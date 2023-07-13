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
  selector: 'app-serial-edit',
  templateUrl: './serial-edit.component.html',
  styleUrls: ['./serial-edit.component.scss'],
})
export class SerialEditComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
  ports: any = [];
  mode = "new";

  constructor(
    private fb: UntypedFormBuilder,
    private msg: NzMessageService,
    private rs: RequestService,
    private route: ActivatedRoute,
    private router: Router
  ) { }

  ngOnInit(): void {
    this.rs.get('serial/ports').subscribe((res) => {
      this.ports = res.data;
    });

    if (this.route.snapshot.paramMap.has('id')) {
      this.mode = "edit";
      this.id = this.route.snapshot.paramMap.get('id');
      this.rs.get(`serial/${this.id}`).subscribe((res) => {
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
      port_name: [mess.port_name || ''],
      poller_period: [mess.poller_period || 60],
      poller_interval: [mess.poller_interval || 2],
      protocol_name: [mess.protocol || 'rtu'],
      protocol_options: [mess.protocol || ''],
      retry_timeout: [mess.retry_timeout || 10],
      retry_maximum: [mess.retry_maximum || 0],
      baud_rate: [mess.baud_rate || 9600],
      parity_mode: [mess.parity_mode || 0],
      stop_bits: [mess.stop_bits || 1],
      data_bits: [mess.data_bits || 8],
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
    this.router.navigateByUrl(`/admin/serial`);
  }

  submit() {
    if (this.validateForm.valid) {
      let url = this.id ? `serial/${this.id}` : `serial/create`;
      const sendData = Object.assign({}, this.validateForm.value);
      sendData.baud_rate = Number(this.validateForm.value.baud_rate)
      this.rs.post(url, sendData).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/serial`);
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
