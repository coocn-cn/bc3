import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { Task } from '../model/task';

@Injectable({
  providedIn: 'root'
})
export class TaskService {

  constructor() { }

  getChilds(parentID: number): Observable<Task[]> {
    let endTime: Date = new Date();
    let endTime2: Date = new Date();

    endTime.setTime(new Date().getTime() + 1000 * 60 * 60 * 24 * 3);
    endTime2.setTime(new Date().getTime() + 1000 * 60 * 60 * 24 * 1);

    let tasks: Task[] = [
      { id: 1, title: "这里是任务标题，这个是待评审需求1", content: "这里是任务详情，这里是为了长度添加的填充语句", createTime: new Date(), complete: false },
      { id: 2, title: "待评审需求，长度要不一样，这个是待评审需求2", content: "这里是任务详情，这里是为了长度添加的填充语句", createTime: new Date(), complete: true },
      { id: 3, title: "待评审需求，这个长度要不一样，这个是待评审需求3", content: "这里是任务详情，这里是为了长度添加的填充语句", createTime: new Date(), complete: false, assignNames: ["聂梦龙"] },
      { id: 4, title: "待评审需求，这个长度确实好像不一样，这个是待评审需求4", content: "这里是任务详情，这里是为了长度添加的填充语句", createTime: new Date(), complete: false, endTime: endTime },
      { id: 5, title: "待评审需求，我说这个不一样，这个是待评审需求5", content: "这里是任务详情，这里是为了长度添加的填充语句", createTime: new Date(), complete: false, assignNames: ["聂梦龙", "高邑爱"], endTime: endTime2 },
    ]

    let groups: Task[] = [
      { id: 1, title: "待评审", content: "", createTime: new Date(), complete: false, childs: tasks },
      { id: 1, title: "已评审", content: "", createTime: new Date(), complete: false, childs: tasks },
      { id: 1, title: "待开发", content: "", createTime: new Date(), complete: false, childs: tasks },
      { id: 1, title: "正在开发", content: "", createTime: new Date(), complete: false, childs: tasks },
      { id: 1, title: "已完成", content: "", createTime: new Date(), complete: false, childs: tasks },
    ]
    return of(groups);
  }
}
