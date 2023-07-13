import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';
import { RequestService } from '../../../request.service';
import { NzMessageService } from 'ng-zorro-antd/message';
@Component({
  selector: 'app-mqtt',
  templateUrl: './mqtt.component.html',
  styleUrls: ['./mqtt.component.scss']
})
export class MqttComponent  implements OnInit {
   
  group!: FormGroup;
  id: any = 0;
  constructor(
    private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private msg: NzMessageService
  ) {this.load()}

  ngOnInit(): void {
    this.group = this.fb.group({
      ClientId: ['', [Validators.required]],
      Username: ['', [Validators.required]],
      Url: ['', [Validators.required]],
      Password: ['', [Validators.required]],
    });
  }
  submit(){
    if (this.group.valid) {
 
      let url = `user/${this.id}` 
      this.rs.post(url, this.group.value).subscribe(res => {
        let path = "/user/list"
        if (location.pathname.startsWith("/admin"))
          path = "/admin" + path
        this.router.navigateByUrl(path)
        this.msg.success("保存成功")
      })
 
     return;
   }
   else {  
    Object.values(this.group.controls).forEach(control => {
      if (control.invalid) {
        control.markAsDirty();
        control.updateValueAndValidity({ onlySelf: true });
      }
    });
     
   }
  }
  loading=false
  query={}
 mqttData=[] 
  load(){ 
      this.rs.get(`config/mqtt`).subscribe((res) => {  
        this.mqttData=res.data
        this.group.patchValue({ClientId:res.data.clientId,Url:res.data.url })
      }); 
  }
   
}
