import { Component, OnInit } from '@angular/core';
import navigation from './navigation.json';

@Component({
  selector: 'framework-navigation',
  templateUrl: './navigation.component.html',
  styleUrls: ['./navigation.component.less']
})
export class NavigationComponent implements OnInit {
  navigation = navigation;
  constructor() { }

  ngOnInit(): void {
  }

}
