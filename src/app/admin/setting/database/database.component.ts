import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from "@angular/forms";
import { ActivatedRoute, Router } from "@angular/router";
import { RequestService } from '../../../request.service';
import { NzMessageService } from "ng-zorro-antd/message";

@Component({
  selector: 'app-database',
  templateUrl: './database.component.html',
  styleUrls: ['./database.component.scss']
})
export class DatabaseComponent implements OnInit {
  group!: FormGroup;
  loading=false
  query={}
  dbData=[] 
  constructor(private fb: FormBuilder,
    private router: Router,
    private route: ActivatedRoute,
    private rs: RequestService,
    private msg: NzMessageService) {this.load()
  }
  switchValue = false;

  ngOnInit(): void {
    // this.rs.get(`config`).subscribe(res => {
    //   //let data = res.data;
    //   this.build(res.data)
    // })

    this.build()
  }

  load(){ 
    this.rs.get(`config/database`).subscribe((res) => {   
     this.dbData=res.data
     this.group.patchValue({LogLevel:String(res.data.log_level)  ,Type:res.data.type,
      URL:res.data.url  })
    }); 
}

  build(obj?: any) {
    obj = obj || {}
    this.group = this.fb.group({
      Type: [obj.type || 'mysql', []],
      URL: ['' ],
      Debug: ['' ],
      LogLevel: ['']
    })
  }

  submit() {
    if (this.group.valid) {
      this.group.patchValue({LogLevel:Number(this.group.value.LogLevel)})
      this.rs.post(`config/database`, this.group.value).subscribe(res => {
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
   
   
}
