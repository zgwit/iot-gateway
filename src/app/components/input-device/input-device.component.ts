import {Component, forwardRef, Input, OnInit} from '@angular/core';
import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/forms";
import {NzInputDirective} from "ng-zorro-antd/input";
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzModalService} from "ng-zorro-antd/modal";
import {SmartRequestService} from "@god-jason/smart";
import {DevicesComponent} from "../../admin/device/devices/devices.component";

@Component({
    selector: 'app-input-device',
    standalone: true,
    imports: [
        NzInputDirective,
        NzButtonComponent
    ],
    templateUrl: './input-device.component.html',
    styleUrl: './input-device.component.scss',
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => InputDeviceComponent),
            multi: true
        }
    ]
})
export class InputDeviceComponent implements OnInit, ControlValueAccessor {
    id = ""
    device: any = {}
    @Input() data: any
    @Input() placeholder = ''
    private onChange!: any;

    constructor(private ms: NzModalService, private rs: SmartRequestService) {
    }

    ngOnInit(): void {
    }

    registerOnChange(fn: any): void {
        this.onChange = fn;
    }

    registerOnTouched(fn: any): void {
    }

    writeValue(obj: any): void {
        if (this.id !== obj) {
            this.id = obj
            if (this.id)
                this.load()
        }
    }

    load() {
        console.log('load device', this.id)
        this.rs.get('device/' + this.id).subscribe(res => {
            if (res.data) {
                this.device = res.data;
            }
        })
    }

    select() {
        this.ms.create({
            nzTitle: "选择",
            nzContent: DevicesComponent,
            nzData: this.data,
        }).afterClose.subscribe(res => {
            console.log(res)
            if (res) {
                this.device = res
                this.id = res.id
                this.onChange(this.id)
            }
        })
    }

    change(value: string) {
        console.log('on change', value)
        this.id = value
        this.onChange(value)
        this.load()
    }
}
