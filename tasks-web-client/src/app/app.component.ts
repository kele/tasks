import { Component } from '@angular/core';
import { CdkDragDrop, moveItemInArray, transferArrayItem } from '@angular/cdk/drag-drop';
import { Client, Service, Task } from "./oto.gen";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  service: Service;

  title = 'tasks-web-client';

  items : { [category: string] : Task[] } = {};

  categories = [
    "TODO",
    "Done",
    "Blocked",
  ];

  constructor() {
    this.service = new Service(new Client());
  }

  ngOnInit() {
    this.categories.forEach(category => {
      this.service.list({category: category}).then(
        allItems => {
          this.service.batchGet(allItems).then(
            tasks => {
              console.log(tasks);
              this.items[category] = tasks.tasks
            }
          )
        }
      )
    });
  }

  drop(event: CdkDragDrop<Task[]>) {
    if (event.previousContainer === event.container) {
      moveItemInArray(event.container.data, event.previousIndex, event.currentIndex);
    } else {
      transferArrayItem(event.previousContainer.data,
        event.container.data,
        event.previousIndex,
        event.currentIndex);
    }
  }
}
