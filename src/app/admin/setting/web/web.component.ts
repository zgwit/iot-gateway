import {Component, OnInit} from '@angular/core';
import {FormBuilder, Validators,FormGroup,} from "@angular/forms";
import {ActivatedRoute, Router} from "@angular/router";
import { RequestService } from '../../../request.service';
import {NzMessageService} from "ng-zorro-antd/message";

@Component({
  selector: 'app-web',
  templateUrl: './web.component.html',
  styleUrls: ['./web.component.scss']
})
export class WebComponent implements OnInit {
  group!: FormGroup;

  constructor(private fb: FormBuilder,
              private router: Router,
              private route: ActivatedRoute,
              private rs: RequestService,
              private msg: NzMessageService) { this.load()
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
    //  web: [obj.web || 8888, []],
    Addr: ['', [Validators.required]],
    Debug: ['', [Validators.required]],
    Cors: ['', [Validators.required]],
    Gzip: ['', [Validators.required]],
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
  webData=[] 
  load(){ 
      this.rs.get(`config/web`).subscribe((res) => {
       
        this.webData=res.data
        this.group.patchValue({Addr:res.data.addr })
      }); 
  }
}
