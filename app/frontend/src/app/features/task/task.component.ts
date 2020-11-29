import { Component, OnInit } from '@angular/core';
import { Task } from './model/task';
import { TaskService } from './service/task.service';

@Component({
  selector: 'app-task',
  templateUrl: './task.component.html',
  styleUrls: ['./task.component.less'],
})
export class TaskComponent implements OnInit {
  tasks: Task[];

  constructor(private taskService: TaskService) { }

  ngOnInit(): void {
    this.taskService.getChilds(0).subscribe(tasks => this.tasks = tasks);
  }

}
