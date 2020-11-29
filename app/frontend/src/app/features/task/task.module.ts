import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { TaskComponent } from './task.component';
import { ItemComponent } from './component/item/item.component';
import { HeaderComponent } from './component/header/header.component';


@NgModule({
  declarations: [
    TaskComponent,
    ItemComponent,
    HeaderComponent,
  ],
  exports:[
    TaskComponent,
  ],
  imports: [
    CommonModule
  ]
})
export class TaskModule { }
