import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'features-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.less']
})
export class HomeComponent implements OnInit {
  title = 'frontend';

  constructor() { }

  ngOnInit(): void {
  }

}
