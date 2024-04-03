import { Component, Input } from '@angular/core';
import { V1Todo } from '../../../../../protos/clients/ts/src/models';
import { CommonModule } from '@angular/common';

@Component({
  selector: '[app-todo-item]',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './todo-item.component.html',
  styleUrl: './todo-item.component.css'
})
export class TodoItemComponent {
  @Input() todo!: V1Todo
}
