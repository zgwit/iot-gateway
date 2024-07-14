import {Component, forwardRef, Input, OnInit} from '@angular/core';
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzInputDirective} from "ng-zorro-antd/input";
import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/forms";
import {NzModalService} from "ng-zorro-antd/modal";
import {ServersComponent} from "../../admin/server/servers/servers.component";
import {SmartRequestService} from "@god-jason/smart";

@Component({
    selector: 'app-input-server',
    standalone: true,
    imports: [
        NzButtonComponent,
        NzInputDirective
    ],
    templateUrl: './input-server.component.html',
    styleUrl: './input-server.component.scss',
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => InputServerComponent),
            multi: true
        }
    ]
})
export class InputServerComponent implements OnInit, ControlValueAccessor {
    id = ""
    server: any = {}
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
        console.log('load servers', this.id)
        this.rs.get('server/' + this.id).subscribe(res => {
            if (res.data) {
                this.server = res.data;
            }
        })
    }

    select() {
        this.ms.create({
            nzTitle: "选择",
            nzContent: ServersComponent,
            nzData: this.data
        }).afterClose.subscribe(res => {
            console.log(res)
            if (res) {
                this.server = res
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
