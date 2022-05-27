# openssl

### 一、使用方法

##### 1、AES的用法
```go
package test

import (
    "github.com/daimayun/go/function"
    "github.com/daimayun/go/openssl"
    "github.com/daimayun/go/openssl/aes"
)

// AesECBEncrypt AES的ECB模式加密
func AesECBEncrypt(data, key string) (string, error) {
    res, err := aes.ECBEncrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return function.Base64Encode(string(res)), nil
}

// AesECBDecrypt AES的ECB模式解密
func AesECBDecrypt(data, key string) (string, error) {
    data = function.Base64Decode(data)
    res, err := aes.ECBDecrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return string(res), nil
}
```

##### 2、DES的用法
```go
package test

import (
    "github.com/daimayun/go/function"
    "github.com/daimayun/go/openssl"
    "github.com/daimayun/go/openssl/des"
)

// DesECBEncrypt DES的ECB模式加密
func DesECBEncrypt(data, key string) (string, error) {
    res, err := des.ECBEncrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return function.Base64Encode(string(res)), nil
}

// DesECBDecrypt DES的ECB模式解密
func DesECBDecrypt(data, key string) (string, error) {
    data = function.Base64Decode(data)
    res, err := des.ECBDecrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return string(res), nil
}
```

##### 3、DES3的用法
```go
package test

import (
    "github.com/daimayun/go/function"
    "github.com/daimayun/go/openssl"
    "github.com/daimayun/go/openssl/des3"
)

// Des3ECBEncrypt DES3的ECB模式加密
func Des3ECBEncrypt(data, key string) (string, error) {
    res, err := des3.ECBEncrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return function.Base64Encode(string(res)), nil
}

// Des3ECBDecrypt DES3的ECB模式解密
func Des3ECBDecrypt(data, key string) (string, error) {
    data = function.Base64Decode(data)
    res, err := des3.ECBDecrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return string(res), nil
}
```

### 二、特别注意

##### 1、实现```php```的加解密方式 ```openssl_encrypt($data, 'DES-EDE', $key)``` 
```go
package test

import (
    "github.com/daimayun/go/function"
    "github.com/daimayun/go/openssl"
    "github.com/daimayun/go/openssl/des3"
)

// Des3ECBEncrypt DES3的ECB模式加密
func Des3ECBEncrypt(data, key string) (string, error) {
    res, err := des3.ECBEncrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return function.Base64Encode(string(res)), nil
}

// Des3ECBDecrypt DES3的ECB模式解密
func Des3ECBDecrypt(data, key string) (string, error) {
    data = function.Base64Decode(data)
    res, err := des3.ECBDecrypt([]byte(data), []byte(key), openssl.PKCS7)
    if err != nil {
        return "", err
    }
    return string(res), nil
}
```
