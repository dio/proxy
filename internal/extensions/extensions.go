// Copyright 2022 Dhi Aurrahman
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package extensions

// TODO(dio): Generate this from *.proto.
import (
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/file/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/filters/cel/v3"                              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/grpc/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/stream/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/access_loggers/wasm/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/cache/simple_http_cache/v3"                                 // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/clusters/aggregate/v3"                                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/clusters/dynamic_forward_proxy/v3"                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/clusters/redis/v3"                                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/dynamic_forward_proxy/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/matching/v3"                                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/ratelimit/v3"                                        // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/common/tap/v3"                                              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/brotli/compressor/v3"                           // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/brotli/decompressor/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/gzip/compressor/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/compression/gzip/decompressor/v3"                           // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/dependency/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/common/matcher/action/v3"                           // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/adaptive_concurrency/v3"                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/admission_control/v3"                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/alternate_protocols_cache/v3"                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/aws_lambda/v3"                                 // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/aws_request_signing/v3"                        // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/bandwidth_limit/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/buffer/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cache/v3"                                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cdn_loop/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/composite/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/compressor/v3"                                 // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/cors/v3"                                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/csrf/v3"                                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/decompressor/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamic_forward_proxy/v3"                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/dynamo/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ext_authz/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ext_proc/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/fault/v3"                                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_http1_bridge/v3"                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_http1_reverse_bridge/v3"                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_json_transcoder/v3"                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_stats/v3"                                 // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/grpc_web/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/gzip/v3"                                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/header_to_metadata/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/health_check/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ip_tagging/v3"                                 // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/jwt_authn/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/kill_request/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/local_ratelimit/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/lua/v3"                                        // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/oauth2/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/on_demand/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/original_src/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/ratelimit/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/rbac/v3"                                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/router/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/set_metadata/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/stateful_session/v3"                           // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/tap/v3"                                        // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/http/wasm/v3"                                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/http_inspector/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_dst/v3"                           // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/original_src/v3"                           // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/proxy_protocol/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/listener/tls_inspector/v3"                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/client_ssl_auth/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/connection_limit/v3"                        // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/direct_response/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/dubbo_proxy/router/v3"                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/dubbo_proxy/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/echo/v3"                                    // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/ext_authz/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/http_connection_manager/v3"                 // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/local_ratelimit/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/meta_protocol_proxy/matcher/action/v3"      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/meta_protocol_proxy/matcher/v3"             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/meta_protocol_proxy/v3"                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/mongo_proxy/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/ratelimit/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/rbac/v3"                                    // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/redis_proxy/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/sni_cluster/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/sni_dynamic_forward_proxy/v3"               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/tcp_proxy/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/filters/header_to_metadata/v3" // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/filters/ratelimit/v3"          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/router/v3"                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/thrift_proxy/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/wasm/v3"                                    // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/network/zookeeper_proxy/v3"                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/udp/dns_filter/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/filters/udp/udp_proxy/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/formatter/metadata/v3"                                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/formatter/req_without_query/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/health_checkers/redis/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/http/header_formatters/preserve_case/v3"                    // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/http/original_ip_detection/custom_header/v3"                // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/http/original_ip_detection/xff/v3"                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/http/stateful_session/cookie/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/internal_redirect/allow_listed_routes/v3"                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/internal_redirect/previous_routes/v3"                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/internal_redirect/safe_cross_scheme/v3"                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/key_value/file_based/v3"                                    // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/load_balancing_policies/round_robin/v3"                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/load_balancing_policies/wrr_locality/v3"                    // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/matching/common_inputs/environment_variable/v3"             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/matching/input_matchers/consistent_hashing/v3"              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/matching/input_matchers/ip/v3"                              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/network/dns_resolver/apple/v3"                              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/network/dns_resolver/cares/v3"                              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/network/socket_interface/v3"                                // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/quic/crypto_stream/v3"                                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/quic/proof_source/v3"                                       // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/rate_limit_descriptors/expr/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/rbac/matchers/upstream_ip_port/v3"                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/request_id/uuid/v3"                                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/resource_monitors/fixed_heap/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/resource_monitors/injected_resource/v3"                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/host/omit_canary_hosts/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/host/omit_host_metadata/v3"                           // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/host/previous_hosts/v3"                               // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/retry/priority/previous_priorities/v3"                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/stat_sinks/graphite_statsd/v3"                              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/stat_sinks/wasm/v3"                                         // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/alts/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/proxy_protocol/v3"                        // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/quic/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/raw_buffer/v3"                            // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/s2a/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/starttls/v3"                              // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tap/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tcp_stats/v3"                             // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/transport_sockets/tls/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/generic/v3"                                  // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/http/v3"                                     // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/tcp/v3"                                      // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/http/v3"                                          // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/upstreams/tcp/generic/v3"                                   // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/wasm/v3"                                                    // to resolve missing type URL
	_ "github.com/envoyproxy/go-control-plane/envoy/extensions/watchdog/profile_action/v3"                                 // to resolve missing type URL
)
