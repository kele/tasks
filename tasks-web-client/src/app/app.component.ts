import { Component } from '@angular/core';
import * as api from "./oto.gen";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  service: api.Service;

  title = 'tasks-web-client';

  items : { [category: string] : api.Task[] } = {};

  categories = [
    "TODO",
    "Done",
    "Blocked",
  ];

  constructor() {
    this.service = new api.Service(new api.Client());
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
}
