package base_interface

// DoModel用途：（写入数据库前Do对象的拦截器）
// 		1、在业务层添加拓展数据 （数据追加）
//		2、数据拦截与置空
// 		3、数据校验
// 使用方式：
// 		1、在业务层赋值 MakeDo
//		2、逻辑实现层调用 DoFactory

type DoModel[TDO interface{}] struct {
	BuildDo func(do TDO) (interface{}, error)
}

// DoFactory 构建待写入数据库的Do数据对象
func (d *DoModel[TDO]) DoFactory(do TDO) (interface{}, error) {
	if d.BuildDo != nil {
		//makeDo, err := d.BuildDo(do)
		//if err != nil {
		//	return makeDo.(TDO), err
		//}

		//tdo, ok := makeDo.(TDO)
		//if !ok {
		//	return makeDo.(TDO), errors.New("do模型不匹配")
		//}
		return d.BuildDo(do)
	}
	return do, nil
}
