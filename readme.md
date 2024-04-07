### 代码复制自gin项目,做了一些使用便利性修改

### 比较规则
| 规则           | 描述               |
| :------------- | ------------------ |
| eq             | 等于               |
| eq_ignore_case | 等于，忽略大小写   |
| ne             | 不等于             |
| ne_ignore_case | 不等于，忽略大小写 |
| gt             | 大于               |
| gte            | 大于等于           |
| lt             | 小于               |
| lte            | 小于等于           |

### 字段规则
| 规则          | 描述                                           |
| ------------- | ---------------------------------------------- |
| eqfield       | 当前字段必须等于指定字段                       |
| nefield       | 当前字段必须不等于指定字段                     |
| gtfield       | 当前字段必须大于指定字段                       |
| gtefield      | 当前字段必须大于等于指定字段                   |
| ltfield       | 当前字段必须小于指定字段                       |
| ltefield      | 当前字段必须小于等于指定字段                   |
| fieldcontains | 当前字段必须包含指定字段，只能用于字符串类型   |
| fieldexcludes | 当前字段必须不包含指定字段，只能用于字符串类型 |
| eqcsfield     | 当前字段必须等于指定结构体的字段               |
| necsfield     | 当前字段必须不等于指定结构体的字段             |
| gtcsfield     | 当前字段必须大于指定结构体的字段               |
| gtecsfield    | 当前字段必须大于等于指定结构体的字段           |
| ltcsfield     | 当前字段必须小于指定结构体的字段               |
| ltecsfield    | 当前字段必须小于等于指定结构体的字段           |

### 字符串规则
| Tag             | Description                                             |
| --------------- | ------------------------------------------------------- |
| alpha           | 仅限字母                                                |
| alphanum        | 仅限字母、数字                                          |
| alphanumunicode | 仅限字母、数字和 Unicode                                |
| alphaunicode    | 字母和 Unicode                                          |
| ascii           | ASCII码字符                                             |
| boolean         | 当前字段必须是能够被 strconv.ParseBool 解析为字符串的值 |
| contains        | 当前字段必须包含指定字符串                              |
| containsany     | Contains Any                                            |
| containsrune    | Contains Rune                                           |
| endsnotwith     | 当前字段必须不是以指定字符串结尾                        |
| endswith        | 当前字段必须以指定字符串结尾                            |
| excludes        | 不包含                                                  |
| excludesall     | Excludes All                                            |
| excludesrune    | Excludes Rune                                           |
| lowercase       | 当前字段的字母必须是小写，可包含数字，不能为空          |
| multibyte       | Multi-Byte Characters                                   |
| number          | 数字                                                    |
| numeric         | Numeric                                                 |
| printascii      | 可打印的ASCII码字符                                     |
| startsnotwith   | 当前字段必须不是以指定字符串开始                        |
| startswith      | 当前字段必须以指定字符串开头                            |
| uppercase       | 当前字段的字母必须是大写，可包含数字，不能为空          |

### 字段规则
| 规则              | 描述                                                         |
| ----------------- | ------------------------------------------------------------ |
| base64            | 当前字段必须是有效的 Base64 值，不能为空                     |
| base64url         | 当前字段必须是包含根据 RFC4648 规范的有效 base64 URL 安全值。 |
| json              | 正确的 JSON 串                                               |
| rgb               | 正确的 RGB 字符串                                            |
| rgba              | 正确的 RGBA 字符串                                           |
| dir               | 指定字段的值必须是已存在的目录                               |
| dirpath           | 指定字段的值必须是合法的目录路径                             |
| file              | 指定字段的值必须是已存在的文件                               |
| filepath          | 指定字段的值必须是合法的文件路径                             |
| len               | 当前字段的长度必须指定值，可用于 string、slice 等            |
| max               | 当前字段的最大值必须是指定值                                 |
| min               | 当前字段的最小值必须是指定值                                 |
| required          | 当前字段为必填项，且不能为零值                               |
| required_if       | 当指定字段等于给定值时，当前字段使用 required 验证。如 required_if=Field1 foo Field2 bar |
| required_unless   | 当指定字段不等于给定值时，当前字段使用 required 验证。如     |
| required_with     | 当任一指定字段不为零值时，当前字段使用 required 验证。如 required_with=Field1 Field2 |
| required_with_all | 当所有指定字段不为零值时，当前字段使用 required 验证。       |

### 网络类规则
| 规则             | 描述                                                         |
| ---------------- | ------------------------------------------------------------ |
| cidr             | 点分格式的ip地址.比如:192.168.1.1或fe80::ba44:827f:4389:6c4b%3 |
| cidrv4           | 点分格式的ipv4地址.比如:192.168.1.1                          |
| cidrv6           | 点分格式的ipv4地址.比如:fe80::ba44:827f:4389:6c4b%3          |
| datauri          | Data URL                                                     |
| fqdn             | Full Qualified Domain Name (FQDN)                            |
| hostname         | 符合RFC 952规范的主机名                                      |
| hostname_port    | 主机的端口号.   0-65535                                      |
| hostname_rfc1123 | 符合 RFC 1123 规范的主机名                                   |
| ip               | Internet Protocol Address IP                                 |
| ip4_addr         | Internet Protocol Address IPv4                               |
| ip6_addr         | Internet Protocol Address IPv6                               |
| ip_addr          | Internet Protocol Address IP                                 |
| ipv4             | Internet Protocol Address IPv4                               |
| ipv6             | Internet Protocol Address IPv6                               |
| mac              | Media Access Control Address MAC                             |
| tcp4_addr        | Transmission Control Protocol Address TCPv4                  |
| tcp6_addr        | Transmission Control Protocol Address TCPv6                  |
| tcp_addr         | Transmission Control Protocol Address TCP                    |
| udp4_addr        | User Datagram Protocol Address UDPv4                         |
| udp6_addr        | User Datagram Protocol Address UDPv6                         |
| udp_addr         | User Datagram Protocol Address UDP                           |
| unix_addr        | Unix domain socket end point Address                         |
| uri              | URI String                                                   |
| url              | URL String                                                   |
| http_url         | HTTP URL String                                              |
| url_encoded      | URL Encoded                                                  |
| urn_rfc2141      | Urn RFC 2141 String                                          |

### 格式化
| Tag                           | Description                                                  |
| ----------------------------- | ------------------------------------------------------------ |
| base64                        | Base64 String                                                |
| base64url                     | Base64URL String                                             |
| base64rawurl                  | Base64RawURL String                                          |
| bic                           | Business Identifier Code (ISO 9362)                          |
| bcp47_language_tag            | Language tag (BCP 47)                                        |
| btc_addr                      | Bitcoin Address                                              |
| btc_addr_bech32               | Bitcoin Bech32 Address (segwit)                              |
| credit_card                   | Credit Card Number                                           |
| mongodb                       | MongoDB ObjectID                                             |
| cron                          | Cron                                                         |
| spicedb                       | SpiceDb ObjectID/Permission/Type                             |
| datetime                      | Datetime                                                     |
| e164                          | e164 formatted phone number                                  |
| email                         | E-mail String                                                |
| eth_addr                      | Ethereum Address                                             |
| hexadecimal                   | Hexadecimal String                                           |
| hexcolor                      | Hexcolor String                                              |
| hsl                           | HSL String                                                   |
| hsla                          | HSLA String                                                  |
| html                          | HTML Tags                                                    |
| html_encoded                  | HTML Encoded                                                 |
| isbn                          | International Standard Book Number                           |
| isbn10                        | International Standard Book Number 10                        |
| isbn13                        | International Standard Book Number 13                        |
| issn                          | International Standard Serial Number                         |
| iso3166_1_alpha2              | Two-letter country code (ISO 3166-1 alpha-2)                 |
| iso3166_1_alpha3              | Three-letter country code (ISO 3166-1 alpha-3)               |
| iso3166_1_alpha_numeric       | Numeric country code (ISO 3166-1 numeric)                    |
| iso3166_2                     | Country subdivision code (ISO 3166-2)                        |
| iso4217                       | Currency code (ISO 4217)                                     |
| json                          | JSON                                                         |
| jwt                           | JSON Web Token (JWT)                                         |
| latitude                      | Latitude                                                     |
| longitude                     | Longitude                                                    |
| luhn_checksum                 | Luhn Algorithm Checksum (for strings and (u)int)             |
| postcode_iso3166_alpha2       | Postcode                                                     |
| postcode_iso3166_alpha2_field | Postcode                                                     |
| rgb                           | RGB String                                                   |
| rgba                          | RGBA String                                                  |
| ssn                           | Social Security Number SSN                                   |
| timezone                      | Timezone                                                     |
| uuid                          | Universally Unique Identifier UUID                           |
| uuid3                         | Universally Unique Identifier UUID v3                        |
| uuid3_rfc4122                 | Universally Unique Identifier UUID v3 RFC4122                |
| uuid4                         | Universally Unique Identifier UUID v4                        |
| uuid4_rfc4122                 | Universally Unique Identifier UUID v4 RFC4122                |
| uuid5                         | Universally Unique Identifier UUID v5                        |
| uuid5_rfc4122                 | Universally Unique Identifier UUID v5 RFC4122                |
| uuid_rfc4122                  | Universally Unique Identifier UUID RFC4122                   |
| md4                           | MD4 hash                                                     |
| md5                           | MD5 hash                                                     |
| sha256                        | SHA256 hash                                                  |
| sha384                        | SHA384 hash                                                  |
| sha512                        | SHA512 hash                                                  |
| ripemd128                     | RIPEMD-128 hash                                              |
| ripemd128                     | RIPEMD-160 hash                                              |
| tiger128                      | TIGER128 hash                                                |
| tiger160                      | TIGER160 hash                                                |
| tiger192                      | TIGER192 hash                                                |
| semver                        | Semantic Versioning 2.0.0                                    |
| ulid                          | Universally Unique Lexicographically Sortable Identifier ULID |
| cve                           | Common Vulnerabilities and Exposures Identifier (CVE id)     |

### 其他
| Tag                  | Description                                                  |
| -------------------- | ------------------------------------------------------------ |
| dir                  | 文件夹是否存在                                               |
| dirpath              | 文件路径                                                     |
| file                 | 文件是否存在                                                 |
| filepath             | 文件路径                                                     |
| image                | Image                                                        |
| isdefault            | Is Default                                                   |
| len                  | 长度:len=10                                                  |
| max                  | 最大值:max=20                                                |
| min                  | 最小值:min=1                                                 |
| oneof                | 枚举值，数值或字符串，以空格分隔，如果字符串中有空格，则使用单引号包围。例如：oneof=beijing shanghai |
| required             | 必须有                                                       |
| required_if          | Required If                                                  |
| required_unless      | Required Unless                                              |
| required_with        | Required With                                                |
| required_with_all    | Required With All                                            |
| required_without     | Required Without                                             |
| required_without_all | Required Without All                                         |
| excluded_if          | Excluded If                                                  |
| excluded_unless      | Excluded Unless                                              |
| excluded_with        | Excluded With                                                |
| excluded_with_all    | Excluded With All                                            |
| excluded_without     | Excluded Without                                             |
| excluded_without_all | Excluded Without All                                         |
| unique               | 唯一值,一般用于slice ,约束slice的成员不重复                  |

### Aliases:
| Tag          | Description                                                  |
| ------------ | ------------------------------------------------------------ |
| iscolor      | 颜色值,支持:hexcolor\|rgb\|rgba\|hsl\|hsla                   |
| country_code | 国家或地区代码:iso3166_1_alpha2\|iso3166_1_alpha3\|iso3166_1_alpha_numeric |

例子:
```go
package main

import (
"fmt"
"github.com/ghp3000/validator"
)

type User struct {
User     string `validate:"min=6,max=32"`
Password string `validate:"min=12,max=32"`
}

func main() {
    err := validator.SetLanguage(validator.LangZh)
    if err != nil {
        fmt.Println(err)
    }
    u := User{User: "admin", Password: "123456"}
    err = validator.ValidateStruct(u)
    if err != nil {
        fmt.Println(err)
    }
}
```