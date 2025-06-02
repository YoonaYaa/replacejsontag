## 此仓库用于将任意对象在输出json时，json中的字段名的tag使用的是自定义的tag，取代原有的默认字段名或json tag

## 例如

type example struct{
<br />    Name string `json:"name" custom:"customName"
<br />}