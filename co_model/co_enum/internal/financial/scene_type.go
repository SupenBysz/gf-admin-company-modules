package financial

import "github.com/kysion/base-library/utility/enum"

// SceneTypeEnum 场景类型：0不限制、1充电佣金收入、
type SceneTypeEnum enum.IEnumCode[int]

type sceneType struct {
    UnLimit SceneTypeEnum
}

var SceneType = sceneType{
    UnLimit: enum.New[SceneTypeEnum](0, "不限制"),
}

func (e sceneType) New(code int, description string) SceneTypeEnum {
    if code == SceneType.UnLimit.Code() {
        return SceneType.UnLimit
    }

    // 可扩展
    return enum.New[SceneTypeEnum](code, description)
}
