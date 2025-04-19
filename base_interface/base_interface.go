package base_interface

// DoModel用途：（写入数据库前Do对象的拦截器）
// 		1、在业务层添加拓展数据 （数据追加）
//		2、数据拦截与置空
// 		3、数据校验
// 使用方式：
// 		1、在业务层赋值 MakeDo, 可以实现业务层DoModel覆盖基础类库的DoModel
//		2、逻辑实现层调用 DoFactory，传入基础类库的DoModel，实现业务层DoModel覆盖基础类库的DoModel
//      3、在业务层赋值 OnSaved 拦截器，在数据保存后执行，返回错误时，数据回滚
//      4、DoSaved 方法，在数据保存后执行，返回错误时，数据会回滚

type DoModel[TDO interface{}] struct {
	BuildDo   func(data TDO) (interface{}, error)
	OnSavedDo func(data TDO, data2 interface{}) error
}

// DoFactory 构建待写入数据库的Do数据对象
func (d *DoModel[TDO]) DoFactory(data TDO) (response interface{}, err error) {
	if d.BuildDo != nil {
		//makeDo, err := d.BuildDo(do)
		//if err != nil {
		//	return makeDo.(TDO), err
		//}

		//tdo, ok := makeDo.(TDO)
		//if !ok {
		//	return makeDo.(TDO), errors.New("do模型不匹配")
		//}
		response, err = d.BuildDo(data)
	}
	return data, err
}

// DoSaved Do数据对象
func (d DoModel[TDO]) DoSaved(data TDO, data2 interface{}) error {
	if d.OnSavedDo != nil {
		return d.OnSavedDo(data, data2)
	}
	return nil
}
