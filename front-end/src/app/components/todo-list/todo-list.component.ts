import { Component, Input } from '@angular/core';
<<<<<<< HEAD
import { TodoItemComponent } from '../todo-item/todo-item.component';
import { CommonModule } from '@angular/common';
=======
import { CommonModule } from '@angular/common';
import { TodoItemComponent } from '../todo-item/todo-item.component';
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a
import { V1Todo } from '../../../../../protos/clients/ts/src/models';

@Component({
  selector: 'app-todo-list',
  standalone: true,
<<<<<<< HEAD
  imports: [CommonModule, TodoItemComponent],
=======
  imports: [TodoItemComponent, CommonModule],
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a
  templateUrl: './todo-list.component.html',
  styleUrl: './todo-list.component.css'
})
export class TodoListComponent {
<<<<<<< HEAD
  @Input() todos: V1Todo[] | undefined;
=======
  @Input() todos: V1Todo[] | undefined
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a
}
