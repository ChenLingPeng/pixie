#include <gtest/gtest.h>
#include <memory>

#include "absl/container/flat_hash_map.h"

#include "src/carnot/compiler/objects/funcobject.h"
#include "src/carnot/compiler/test_utils.h"

namespace pl {
namespace carnot {
namespace compiler {
using ::testing::ElementsAre;
class PyFuncTest : public OperatorTests {};

class TestQLObject : public QLObject {
 public:
  static constexpr TypeDescriptor TestQLObjectType = {
      /* name */ "TestQLObject",
      /* type */ QLObjectType::kMisc,
  };
  explicit TestQLObject(int64_t value) : QLObject(TestQLObjectType), value_(value) {}

  // Value used for testing that we can initialize QLObjects within a FuncObject
  int64_t value() { return value_; }
  void AddKwargValue(const std::string& arg_name, const std::string& value) {
    kwarg_names_.push_back(arg_name);
    kwarg_values_.push_back(value);
  }
  const std::vector<std::string>& kwarg_names() const { return kwarg_names_; }
  const std::vector<std::string>& kwarg_values() const { return kwarg_values_; }

 private:
  int64_t value_;
  std::vector<std::string> kwarg_names_;
  std::vector<std::string> kwarg_values_;
};

StatusOr<QLObjectPtr> SimpleFunc(const pypa::AstPtr&, const ParsedArgs& args) {
  IRNode* node = args.GetArg("simple");
  if (!Match(node, Int())) {
    return node->CreateIRNodeError("Expected int");
  }
  auto out_obj = std::make_shared<TestQLObject>(static_cast<IntIR*>(node)->val());
  for (const auto& [arg, value] : args.kwargs()) {
    if (!Match(value, String())) {
      return value->CreateIRNodeError("Expected string $0", value->DebugString());
    }
    out_obj->AddKwargValue(arg, static_cast<StringIR*>(value)->str());
  }

  return StatusOr<QLObjectPtr>(out_obj);
}

TEST_F(PyFuncTest, PosArgsExecute) {
  std::shared_ptr<FuncObject> func_obj(
      new FuncObject("func", {"simple"}, {}, /*has_kwargs*/ false,
                     std::bind(&SimpleFunc, std::placeholders::_1, std::placeholders::_2)));

  ArgMap args{{}, {MakeInt(123)}};
  std::shared_ptr<QLObject> obj = func_obj->Call(args, ast).ConsumeValueOrDie();
  ASSERT_TRUE(obj->type_descriptor().type() == QLObjectType::kMisc);
  auto test_obj = static_cast<TestQLObject*>(obj.get());
  EXPECT_EQ(test_obj->value(), 123);

  EXPECT_THAT(test_obj->kwarg_names(), ElementsAre());
  EXPECT_THAT(test_obj->kwarg_values(), ElementsAre());
}

TEST_F(PyFuncTest, MissingArgument) {
  std::shared_ptr<FuncObject> func_obj(
      new FuncObject("func", {"simple"}, {}, /*has_kwargs*/ false,
                     std::bind(&SimpleFunc, std::placeholders::_1, std::placeholders::_2)));

  ArgMap args;
  auto status = func_obj->Call(args, ast);
  ASSERT_NOT_OK(status);
  EXPECT_THAT(status.status(),
              HasCompilerError("func.* missing 1 required positional arguments 'simple'"));
}

TEST_F(PyFuncTest, ExtraPosArg) {
  std::shared_ptr<FuncObject> func_obj(
      new FuncObject("func", {"simple"}, {}, /*has_kwargs*/ false,
                     std::bind(&SimpleFunc, std::placeholders::_1, std::placeholders::_2)));

  ArgMap args{{}, {MakeInt(1), MakeInt(2)}};
  auto status = func_obj->Call(args, ast);
  ASSERT_NOT_OK(status);
  EXPECT_THAT(status.status(), HasCompilerError("func.* takes 1 arguments but 2 were given."));

  // Should do the same with kwarg support enabled.
  std::shared_ptr<FuncObject> func_obj2(
      new FuncObject("func", {"simple"}, {}, /*has_kwargs*/ true,
                     std::bind(&SimpleFunc, std::placeholders::_1, std::placeholders::_2)));

  auto status2 = func_obj2->Call(args, ast);
  ASSERT_NOT_OK(status2);
  EXPECT_THAT(status2.status(), HasCompilerError("func.* takes 1 arguments but 2 were given."));
}

TEST_F(PyFuncTest, ExtraKwargNoKwargsSupport) {
  std::shared_ptr<FuncObject> func_obj(
      new FuncObject("func", {"simple"}, {}, /*has_kwargs*/ false,
                     std::bind(&SimpleFunc, std::placeholders::_1, std::placeholders::_2)));

  ArgMap args{{{"blah", MakeString("blah")}}, {MakeInt(1)}};
  auto status = func_obj->Call(args, ast);
  ASSERT_NOT_OK(status);
  EXPECT_THAT(status.status(),
              HasCompilerError("func.* got an unexpected keyword argument 'blah'"));
}

TEST_F(PyFuncTest, ExtraKwargWithKwargsSupport) {
  std::shared_ptr<FuncObject> func_obj(
      new FuncObject("func", {"simple"}, {}, /*has_kwargs*/ true,
                     std::bind(&SimpleFunc, std::placeholders::_1, std::placeholders::_2)));

  ArgMap args{{{"blah1", MakeString("blah2")}}, {MakeInt(123)}};
  std::shared_ptr<QLObject> obj = func_obj->Call(args, ast).ConsumeValueOrDie();
  ASSERT_TRUE(obj->type_descriptor().type() == QLObjectType::kMisc);
  auto test_obj = static_cast<TestQLObject*>(obj.get());
  EXPECT_EQ(test_obj->value(), 123);

  EXPECT_THAT(test_obj->kwarg_names(), ElementsAre("blah1"));
  EXPECT_THAT(test_obj->kwarg_values(), ElementsAre("blah2"));
}

// TODO(philkuz) (PL-1129) figure out default arguments.
TEST_F(PyFuncTest, DISABLED_DefaultArgsExecute) {
  std::shared_ptr<FuncObject> func_obj(
      new FuncObject("func", {"simple"}, {{"simple", "1234"}}, /*has_kwargs*/ false,
                     std::bind(&SimpleFunc, std::placeholders::_1, std::placeholders::_2)));

  ArgMap args;
  std::shared_ptr<QLObject> obj = func_obj->Call(args, ast).ConsumeValueOrDie();
  ASSERT_TRUE(obj->type_descriptor().type() == QLObjectType::kMisc);
  auto test_obj = static_cast<TestQLObject*>(obj.get());
  EXPECT_EQ(test_obj->value(), 1234);
}

}  // namespace compiler
}  // namespace carnot
}  // namespace pl
