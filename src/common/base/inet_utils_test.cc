#include "src/common/base/inet_utils.h"
#include "src/common/testing/testing.h"

namespace pl {

TEST(ParseSockAddr, IPv4) {
  // Create an IP address for the test.
  struct sockaddr_in sockaddr;
  sockaddr.sin_family = AF_INET;
  inet_pton(AF_INET, "10.1.2.3", &sockaddr.sin_addr);
  sockaddr.sin_port = htons(53000);

  // Now check InetAddrToString produces the expected string.
  IPAddress addr;
  Status s = ParseSockAddr(reinterpret_cast<struct sockaddr*>(&sockaddr), &addr);
  EXPECT_OK(s);
  EXPECT_EQ(addr.addr_str, "10.1.2.3");
  EXPECT_EQ(addr.port, 53000);
}

TEST(ParseSockAddr, IPv6) {
  struct sockaddr_in6 sockaddr;
  sockaddr.sin6_family = AF_INET6;
  EXPECT_OK(ParseIPv6Addr("::1", &sockaddr.sin6_addr));
  sockaddr.sin6_port = htons(12345);

  IPAddress addr;
  Status s = ParseSockAddr(reinterpret_cast<struct sockaddr*>(&sockaddr), &addr);
  EXPECT_OK(s);
  EXPECT_EQ(addr.addr_str, "::1");
  EXPECT_EQ(addr.port, 12345);
}

TEST(ParseSockAddr, Unsupported) {
  // Create an IP address for the test.
  struct sockaddr_in sockaddr;
  sockaddr.sin_family = AF_UNIX;
  inet_pton(AF_INET, "10.1.2.3", &sockaddr.sin_addr);
  sockaddr.sin_port = htons(53000);

  IPAddress addr;
  Status s = ParseSockAddr(reinterpret_cast<struct sockaddr*>(&sockaddr), &addr);
  EXPECT_NOT_OK(s);
  EXPECT_EQ(addr.addr_str, "-") << "addr_str should not be mutated";
  EXPECT_EQ(addr.port, -1) << "port should not be mutated";
}

TEST(ParseIPAddr, ipv4) {
  // Test address.
  struct in_addr in_addr;
  EXPECT_OK(ParseIPv4Addr("1.2.3.4", &in_addr));

  // Now check for the expected string.
  std::string addr;
  Status s = ParseIPv4Addr(in_addr, &addr);
  EXPECT_OK(s);
  EXPECT_EQ(addr, "1.2.3.4");
}

TEST(ParseIPAddr, ipv6) {
  // Test address.
  struct in6_addr in6_addr;
  EXPECT_OK(ParseIPv6Addr("2001:0db8:85a3:0000:0000:8a2e:0370:7334", &in6_addr));

  // Now check for the expected string.
  std::string addr;
  Status s = ParseIPv6Addr(in6_addr, &addr);
  EXPECT_OK(s);
  // Note that formatting is slightly different (zeros removed).
  EXPECT_EQ(addr, "2001:db8:85a3::8a2e:370:7334");
}

TEST(ParseIPAddr, ipv4_using_in6_addr) {
  // Create an IP address for the test.
  struct in6_addr in6_addr;
  EXPECT_OK(ParseIPv4Addr("1.2.3.4", &in6_addr));

  // Now check for the expected string.
  std::string addr;
  Status s = ParseIPv4Addr(in6_addr, &addr);
  EXPECT_OK(s);
  // Note that formatting is slightly different (zeros removed).
  EXPECT_EQ(addr, "1.2.3.4");
}

}  // namespace pl
