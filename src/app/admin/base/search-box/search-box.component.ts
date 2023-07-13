import { Component, Input, Output, EventEmitter } from '@angular/core';
@Component({
  selector: 'app-search-box',
  templateUrl: './search-box.component.html',
  styleUrls: ['./search-box.component.scss']
})
export class SearchBoxComponent {
  text = "";
  @Input() placeholder = "关键字";
  @Output() onSearch = new EventEmitter<string>();
}
