import { Time } from '@angular/common';

export interface Task {
    id: number,
    title: string,
    content: string,
    complete: boolean,
    createTime: Date,
    endTime?: Date,
    startTime?: Date,
    manHour?: number,
    parentID?: number,
    childs?: Task[],
    groupName?: string,
    assignNames?: string[],
}
