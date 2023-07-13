import { RequestService } from '../../../request.service';
import { Component, OnInit } from '@angular/core';
import {
  UntypedFormBuilder,
  FormGroup,
  Validators,
} from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-client-edit',
  templateUrl: './client-edit.component.html',
  styleUrls: ['./client-edit.component.scss'],
})
export class ClientEditComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
  mode: string = 'new';
  constructor(
    private fb: UntypedFormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) { }
  ngOnInit(): void {
    if (this.route.snapshot.paramMap.has('id')) {
      this.id = this.route.snapshot.paramMap.get('id');
      this.mode = 'edit';
      this.rs.get(`client/${this.id}`).subscribe((res) => {
        this.setData(res);
      });
    }

    this.build();
  }

  build(mess?: any) {
    mess = mess || {};
    this.validateForm = this.fb.group({
      id: [mess.id || '', this.mode === "edit" ? [Validators.required] : ''],
      name: [mess.name || ''],
      net: [mess.net || 'tcp'],
      addr: [mess.addr || ''],
      port: [mess.port || 1],
      poller_period: [mess.poller_period || 60],
      poller_interval: [mess.poller_interval || 2],
      protocol_name: [mess.protocol || 'rtu'],
      protocol_options: [mess.protocol || 'rtu'],
      retry_timeout: [mess.retry_timeout || 10],
      retry_maximum: [mess.retry_maximum || 0],
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
  }
  handleCancel() {
    this.router.navigateByUrl(`/admin/client`);
  }
  submit() {
    if (this.validateForm.valid) {
      let url = this.id ? `client/${this.id}` : `client/create`;
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/client`);
      });
    } else {
      Object.values(this.validateForm.controls).forEach((control) => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }
}
