#include "src/common/testing/grpc_utils/grpc_utils.h"

namespace pl {
// The following are copied from
// https://github.com/grpc/grpc/blob/v1.29.1/test/core/util/test_config.cc

bool BuiltUnderValgrind() {
#ifdef RUNNING_ON_VALGRIND
  return true;
#else
  return false;
#endif
}

bool BuiltUnderTsan() {
#if defined(__has_feature)
#if __has_feature(thread_sanitizer)
  return true;
#else
  return false;
#endif
#else
#ifdef THREAD_SANITIZER
  return true;
#else
  return false;
#endif
#endif
}

bool BuiltUnderAsan() {
#if defined(__has_feature)
#if __has_feature(address_sanitizer)
  return true;
#else
  return false;
#endif
#else
#ifdef ADDRESS_SANITIZER
  return true;
#else
  return false;
#endif
#endif
}

bool BuiltUnderMsan() {
#if defined(__has_feature)
#if __has_feature(memory_sanitizer)
  return true;
#else
  return false;
#endif
#else
#ifdef MEMORY_SANITIZER
  return true;
#else
  return false;
#endif
#endif
}

bool BuiltUnderUbsan() {
#ifdef GRPC_UBSAN
  return true;
#else
  return false;
#endif
}

int64_t grpc_test_sanitizer_slowdown_factor() {
  int64_t sanitizer_multiplier = 1;
  if (BuiltUnderValgrind()) {
    sanitizer_multiplier = 20;
  } else if (BuiltUnderTsan()) {
    sanitizer_multiplier = 5;
  } else if (BuiltUnderAsan()) {
    sanitizer_multiplier = 3;
  } else if (BuiltUnderMsan()) {
    sanitizer_multiplier = 4;
  } else if (BuiltUnderUbsan()) {
    sanitizer_multiplier = 5;
  }
  return sanitizer_multiplier;
}

/**
 * @brief Sleeps for specified period of time, slowing down based on the build configuration (ie
 * ASAN, TSAN, etc)
 * @param ms_to_wait
 */
void TestSleep(int64_t ms_to_wait) {
  gpr_sleep_until(gpr_time_add(gpr_now(GPR_CLOCK_MONOTONIC),
                               gpr_time_from_micros(grpc_test_sanitizer_slowdown_factor() *
                                                        static_cast<int64_t>(1e3) * ms_to_wait,
                                                    GPR_TIMESPAN)));
}

}  // namespace pl
