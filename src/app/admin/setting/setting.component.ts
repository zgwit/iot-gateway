import {Component, ViewChild} from '@angular/core';
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzCardComponent} from "ng-zorro-antd/card";
import {SmartEditorComponent, SmartField, SmartRequestService} from "@god-jason/smart";
import {NzMessageService} from "ng-zorro-antd/message";
import {ActivatedRoute} from "@angular/router";

@Component({
    selector: 'app-setting',
    standalone: true,
    imports: [
        NzButtonComponent,
        NzCardComponent,
        SmartEditorComponent
    ],
    templateUrl: './setting.component.html',
    styleUrl: './setting.component.scss'
})
export class SettingComponent {
    module: any = "web"

    @ViewChild('form') form!: SmartEditorComponent

    fields: SmartField[] = []

    data: any = {}

    constructor(private msg: NzMessageService,
                private rs: SmartRequestService,
                private route: ActivatedRoute,
    ) {
        //this.module = this.route.snapshot.queryParamMap.get("module")
        this.route.queryParamMap.subscribe(res => {
            this.module = res.get("module")
            this.load()
        })
    }

    load(): void {
        this.rs.get('setting/' + this.module + '/form').subscribe(res => {
            this.fields = res.data
        });
        this.rs.get('setting/' + this.module).subscribe(res => {
            this.data = res.data
        });
    }

    onSubmit() {
        if (!this.form.valid) {
            this.msg.error('请检查数据')
            return
        }

        this.rs.post('setting/' + this.module, this.form.value).subscribe((res) => {
            this.msg.success('保存成功');
        });
    }
}
