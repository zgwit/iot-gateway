import {Component, forwardRef, Input, OnInit} from '@angular/core';
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzInputDirective} from "ng-zorro-antd/input";
import {ControlValueAccessor, NG_VALUE_ACCESSOR} from "@angular/forms";
import {NzModalService} from "ng-zorro-antd/modal";
import {SmartRequestService} from "@god-jason/smart";
import {TunnelsComponent} from "../../admin/tunnel/tunnels/tunnels.component";

@Component({
    selector: 'app-input-tunnel',
    standalone: true,
    imports: [
        NzButtonComponent,
        NzInputDirective
    ],
    templateUrl: './input-tunnel.component.html',
    styleUrl: './input-tunnel.component.scss',
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => InputTunnelComponent),
            multi: true
        }
    ]
})
export class InputTunnelComponent implements OnInit, ControlValueAccessor {
    id = ""
    tunnel: any = {}
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
        }
    }

    select() {
        this.ms.create({
            nzTitle: "选择", nzContent: TunnelsComponent, nzData: this.data
        }).afterClose.subscribe(res => {
            console.log(res)
            if (res) {
                this.tunnel = res
                this.id = res.id
                this.onChange(this.id)
            }
        })
    }

    change(value: string) {
        this.id = value
        this.onChange(value)
    }
}
