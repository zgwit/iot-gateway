package device

import (
	"errors"
	"github.com/god-jason/bucket/pkg/event"
	"github.com/zgwit/iot-gateway/db"
	"github.com/zgwit/iot-gateway/protocol"
	"time"
)

func init() {
	db.Register(new(Device))
}

type Device struct {
	Id string `json:"id" xorm:"pk"` //ClientID

	ProductId string `json:"product_id,omitempty" xorm:"index"`
	Product   string `json:"product,omitempty" xorm:"<-"`

	ProjectId string `json:"project_id,omitempty" xorm:"index"`
	Project   string `json:"project,omitempty" xorm:"<-"`

	Name        string         `json:"name"`
	Description string         `json:"description,omitempty"`
	Parameters  map[string]any `json:"parameters,omitempty" xorm:"json"` //模型参数，用于报警检查
	Disabled    bool           `json:"disabled,omitempty"`
	Created     time.Time      `json:"created,omitempty" xorm:"created"`

	Online bool `json:"online,omitempty" xorm:"-"`

	//通道ID
	TunnelId string         `json:"tunnel_id,omitempty" xorm:"index"`
	Station  map[string]any `json:"station,omitempty" xorm:"json"` //通道参数 保存从站号等

	//变量
	values map[string]any
	//last   time.Time

	//事件监听
	eventData event.Emitter[map[string]any]

	adapter protocol.Adapter
}

func (d *Device) Watch(fn func(map[string]any)) int {
	return d.eventData.On(fn)
}

func (d *Device) UnWatch(handler int) {
	d.eventData.Off(handler)
}

func (d *Device) Values() map[string]any {
	return d.values
}

func (d *Device) Push(values map[string]any) {

	//广播
	d.eventData.Emit(values)

	//赋值
	if d.values == nil {
		d.values = make(map[string]any)
	}
	for k, v := range values {
		d.values[k] = v
	}

	//检查数据
	//d.Validate()
}

func (d *Device) Write(point string, value any) error {
	if d.adapter == nil {
		return errors.New("未连接")
	}
	return d.adapter.Set(d.Id, point, value)
}

func (d *Device) WriteMany(values map[string]any) error {
	if d.adapter == nil {
		return errors.New("未连接")
	}
	for point, value := range values {
		err := d.adapter.Set(d.Id, point, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Device) SetAdapter(adapter protocol.Adapter) {
	d.adapter = adapter //TODO 会内存泄露，需要手动清空
}

//func (d *Device) Validate() {
//	for _, v := range d.validators {
//		ret := v.Validate(d.values)
//		if !ret {
//			//检查结果为真时，才产生报警
//			continue
//		}
//
//		//入库
//		al := alarm.Alarm{
//			ProductId: d.ProductId,
//			Product:   d.Product,
//			DeviceId:  d.Id,
//			Device:    d.Name,
//			Type:      v.Type,
//			Title:     v.Title,
//			Level:     v.Level,
//			Message:   v.Template, //TODO 模板格式化
//		}
//		_, err := db.Engine.Insert(&al)
//		if err != nil {
//			log.Error(err)
//			//continue
//		}
//
//		//通知
//		//err = internal.notify(&al)
//		//if err != nil {
//		//	log.Error(err)
//		//	//continue
//		//}
//	}
//}
