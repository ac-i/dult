# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [proto/speedpb/speed.proto](#proto_speedpb_speed-proto)
    - [Options](#proto-speedpb-Options)
    - [Speed](#proto-speedpb-Speed)
    - [TestResult](#proto-speedpb-TestResult)
  
    - [Provider](#proto-speedpb-Provider)
  
    - [SpeedService](#proto-speedpb-SpeedService)
  
- [Scalar Value Types](#scalar-value-types)



<a name="proto_speedpb_speed-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## proto/speedpb/speed.proto



<a name="proto-speedpb-Options"></a>

### Options
Options


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| provider | [Provider](#proto-speedpb-Provider) |  | provider |






<a name="proto-speedpb-Speed"></a>

### Speed
Speed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| ip | [string](#string) |  | ip |
| host | [string](#string) |  | host |
| download_mbps | [double](#double) |  | download_mbps in Mbps (float64) |
| upload_mbps | [double](#double) |  | upload_mbps in Mbps (float64) |
| latency_ns | [int64](#int64) |  | latency_ns in ns |






<a name="proto-speedpb-TestResult"></a>

### TestResult
TestResult


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| speed | [Speed](#proto-speedpb-Speed) |  | speed |





 


<a name="proto-speedpb-Provider"></a>

### Provider
Provider

| Name | Number | Description |
| ---- | ------ | ----------- |
| PROVIDER_UNSPECIFIED | 0 | PROVIDER_UNSPECIFIED |
| PROVIDER_SPEEDTEST_NET | 1 | PROVIDER_SPEEDTEST_NET |
| PROVIDER_FAST_COM | 2 | PROVIDER_FAST_COM |


 

 


<a name="proto-speedpb-SpeedService"></a>

### SpeedService
SpeedService

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| QuickDULT | [Options](#proto-speedpb-Options) | [TestResult](#proto-speedpb-TestResult) | QuickDULT |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

