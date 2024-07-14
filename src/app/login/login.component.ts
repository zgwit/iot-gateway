import {Component, OnInit} from '@angular/core';
import {Router} from '@angular/router';
import {FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {Md5} from 'ts-md5';
import {NzFormControlComponent, NzFormDirective, NzFormItemComponent} from "ng-zorro-antd/form";
import {NzInputDirective, NzInputGroupComponent} from "ng-zorro-antd/input";
import {NzColDirective, NzRowDirective} from "ng-zorro-antd/grid";
import {NzCheckboxComponent} from "ng-zorro-antd/checkbox";
import {NzButtonComponent} from "ng-zorro-antd/button";
import {NzCardComponent} from "ng-zorro-antd/card";
import {SmartRequestService} from "@god-jason/smart";

@Component({
    selector: 'app-login',
    standalone: true,
    imports: [
        NzFormDirective,
        NzFormItemComponent,
        NzInputGroupComponent,
        NzFormControlComponent,
        NzRowDirective,
        NzColDirective,
        NzCheckboxComponent,
        NzButtonComponent,
        FormsModule,
        ReactiveFormsModule,
        NzInputDirective,
        NzCardComponent
    ],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss'
})
export class LoginComponent implements OnInit {
    validateForm!: FormGroup;

    constructor(private fb: FormBuilder,
                private rs: SmartRequestService,
                private router: Router,
    ) {
    }


    submitForm(): void {
        for (const i in this.validateForm.controls) {
            this.validateForm.controls[i].markAsDirty();
            this.validateForm.controls[i].updateValueAndValidity();
        }
        if (!this.validateForm.valid) {
            return;
        }

        const password = Md5.hashStr(this.validateForm.value.password);

        this.rs.post('login', {username: this.validateForm.value.username, password}).subscribe(res => {
            console.log('res:', res);
            //localStorage.setItem('token', res.data.token);

            //更新用户
            //this.us.setUser(res.data);

            if (res.data.admin) {
                this.router.navigateByUrl('/admin')
            } else {
                this.router.navigateByUrl('/select')
            }
        });
    }

    ngOnInit(): void {
        this.validateForm = this.fb.group({
            username: [null, [Validators.required]],
            password: [null, [Validators.required]],
            remember: [false]
        });
    }
}
