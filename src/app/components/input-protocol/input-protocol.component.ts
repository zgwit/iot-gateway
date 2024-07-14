import {Component, forwardRef, OnInit} from '@angular/core';
import {NzSelectComponent} from "ng-zorro-antd/select";
import {ControlValueAccessor, FormsModule, NG_VALUE_ACCESSOR} from "@angular/forms";
import {SmartRequestService} from "@god-jason/smart";

@Component({
    selector: 'app-input-protocol',
    standalone: true,
    imports: [
        NzSelectComponent,
        FormsModule
    ],
    templateUrl: './input-protocol.component.html',
    styleUrl: './input-protocol.component.scss',
    providers: [
        {
            provide: NG_VALUE_ACCESSOR,
            useExisting: forwardRef(() => InputProtocolComponent),
            multi: true
        }
    ]
})
export class InputProtocolComponent implements OnInit, ControlValueAccessor {

    options: any[] = []
    private onChange!: any;

    constructor(private rs: SmartRequestService) {
        this.load()
    }

    _value: string = ""

    get value() {
        return this._value
    }

    set value(v) {
        this._value = v
        this.onChange(v)
    }

    ngOnInit(): void {
    }

    load() {
        this.rs.get(`protocol/list`).subscribe((res) => {
            this.options = res.data.map((p: any) => {
                return {value: p.name, label: p.label}
            })
        });
    }

    registerOnChange(fn: any): void {
        this.onChange = fn
    }

    registerOnTouched(fn: any): void {
    }

    setDisabledState(isDisabled: boolean): void {
    }

    writeValue(obj: any): void {
        this._value = obj
    }
}
