import { Component, Input } from '@angular/core';
import { V1Todo } from '../../../../../protos/clients/ts/src/models';
<<<<<<< HEAD
import { CommonModule } from '@angular/common';
=======
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a

@Component({
  selector: '[app-todo-item]',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './todo-item.component.html',
  styleUrl: './todo-item.component.css'
})
export class TodoItemComponent {
<<<<<<< HEAD
  @Input() todo!: V1Todo
=======
  @Input() todo!: V1Todo;
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a
}
