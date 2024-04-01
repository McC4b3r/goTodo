import { Component, Input } from '@angular/core';
import { V1Todo } from '../../../../../protos/clients/ts/src/models';

@Component({
  selector: 'app-todo-item',
  standalone: true,
  imports: [],
  templateUrl: './todo-item.component.html',
  styleUrl: './todo-item.component.css'
})
export class TodoItemComponent {
  @Input() todo!: V1Todo;
}
