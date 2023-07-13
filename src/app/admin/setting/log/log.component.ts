import {Component, OnInit} from '@angular/core';
import {FormBuilder, FormGroup, Validators} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import { RequestService } from '../../../request.service';
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-log',
  templateUrl: './log.component.html',
  styleUrls: ['./log.component.scss']
})
export class LogComponent implements OnInit {
  group!: FormGroup;

  constructor(private fb: FormBuilder,
              private router: Router,
              private route: ActivatedRoute,
              private rs: RequestService,
              private msg: NzMessageService) {this.load()
  }


  ngOnInit(): void {
    // this.rs.get(`config`).subscribe(res => {
    //   //let data = res.data;
    //   this.build(res.data)
    // })

    this.build()
  }

  build(obj?: any) {
    obj = obj || {}
    this.group = this.fb.group({ 
      Filename: ['', [Validators.required]],
      MaxSize: ['', [Validators.required]],
      MaxAge: ['', [Validators.required]] ,
      MaxBackups: ['', [Validators.required]],
      Level: ['', [Validators.required]],
      Text: ['', [Validators.required]],
      Format: ['', [Validators.required]] ,
      Output: ['', [Validators.required]],
      Caller: ['', [Validators.required]]
    })
  }

  submit() {
    
    if (this.group.valid) {
 
      this.rs.post(`config`, this.group.value).subscribe(res => {
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
  logData=[] 
  load(){ 
      this.rs.get(`config/log`).subscribe((res) => {    
        this.logData=res.data 
        this.group.patchValue({Caller:res.data.caller,Level:String(res.data.level)   ,Text:res.data.text,
          Output:JSON.stringify(res.data.output)  })
      }); 
  }
}
