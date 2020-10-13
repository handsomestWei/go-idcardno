# go-idcardno
golang 身份证号码工具库。提供18位身份证号码自动生成、有效性校验、信息解析

# Usage
身份证号码自动生成
```
idcardno.AutoCreate18IdCardNo()
```

身份证号码有效性校验
```
idNo := "xxx"
idcardno.Validate18IdCardNo(idNo)
```

身份证号码信息解析，包含行政区域名称、生日、年龄、性别
```
idNo := "xxx"
result, info := idcardno.Parse18IdCardNoInfo(idNo)
if result {
    fmt.Println(fmt.Sprintf("%v", *info))
}
```
