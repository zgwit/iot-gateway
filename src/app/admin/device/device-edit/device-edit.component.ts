import { RequestService } from '../../../request.service';
import { Component, OnInit, ViewChild } from '@angular/core';
import {
  UntypedFormBuilder,
  FormGroup,
  Validators,
} from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ActivatedRoute, Router } from '@angular/router';
@Component({
  selector: 'app-device-edit',
  templateUrl: './device-edit.component.html',
  styleUrls: ['./device-edit.component.scss'],
})
export class DeviceEditComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
  mode = "new";
  tunnel_id: string = '';
  @ViewChild('setProductIdTag') setProductIdTag: any;
  constructor(
    private fb: UntypedFormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) { }
  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.mode = "edit";
      this.id = this.route.snapshot.paramMap.get('id');
      this.rs.get(`device/${this.id}`).subscribe((res) => {
        this.setData(res);
      });
    } else {
      this.route.queryParams.subscribe((params) => {
        this.tunnel_id = params['tunnelId'];
      })
    }
    this.build();
  }
  build(mess?: any) {
    mess = mess || {};
    this.validateForm = this.fb.group({
      id: [mess.id || '', this.mode === "edit" ? [Validators.required] : ''],
      name: [mess.name || ''],
      desc: [mess.desc || ''],
      tunnel_id: [mess.tunnel_id || this.tunnel_id],
      product_id: [mess.product_id || ''],
      slave: [mess.slave || 1]
    });
  }
  setData(res: any) {
    const resData = (res && res.data) || {};
    const odata = this.validateForm.value;
    for (const key in odata) {
      if (resData[key]) {
        odata[key] = resData[key];
      }
    }
    this.validateForm.setValue(odata);
    // 给子组件设值
    this.setProductIdTag.product_id = resData['product_id'] || '';
  }
  handleCancel() {
    this.router.navigateByUrl(`/admin/device`);
  }
  submit() {
    if (this.validateForm.valid) {
      let url = this.id ? `device/${this.id}` : `device/create`;
      this.validateForm.patchValue({ product_id: this.setProductIdTag.product_id })
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/device`);
      });
      return;
    } else {
      Object.values(this.validateForm.controls).forEach((control) => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
  reset() {
    this.validateForm.reset();
    for (const key in this.validateForm.controls) {
      if (this.validateForm.controls.hasOwnProperty(key)) {
        this.validateForm.controls[key].markAsPristine();
        this.validateForm.controls[key].updateValueAndValidity();
      }
    }
  }
}
