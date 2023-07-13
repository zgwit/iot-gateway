import { RequestService } from '../../../request.service';
import { Component, OnInit } from '@angular/core';
import {
  UntypedFormBuilder,
  FormGroup,
} from '@angular/forms';
import { NzMessageService } from 'ng-zorro-antd/message';
import { ActivatedRoute, Router } from '@angular/router';

@Component({
  selector: 'app-link-edit',
  templateUrl: './link-edit.component.html',
  styleUrls: ['./link-edit.component.scss'],
})
export class LinkEditComponent implements OnInit {
  validateForm!: FormGroup;
  id: any = 0;
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
      this.rs.get(`link/${this.id}`).subscribe((res) => {
        this.patchValue(res.data);
      });
    }

    this.patchValue();
  }
  handleCancel() {
    this.router.navigateByUrl(`/admin/link`);
  }

  patchValue(mess?: any) {
    mess = mess || {};
    this.validateForm = this.fb.group({
      id: [mess.id || ''],
      name: [mess.name || ''],
      poller_period: [mess.poller_period || 60],
      poller_interval: [mess.poller_interval || 2],
      protocol_name: [mess.protocol || 'rtu'],
      protocol_options: [mess.protocol || 'rtu'],
    });
  }
  submit() {
    if (this.validateForm.valid) {
      let url = this.id ? `link/${this.id}` : `link/create`;
      this.rs.post(url, this.validateForm.value).subscribe((res) => {
        this.msg.success('保存成功');
        this.router.navigateByUrl(`/admin/link`);
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
