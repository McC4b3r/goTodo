import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { TodoListComponent } from './components/todo-list/todo-list.component';
import { TodoApiResponse, TodoService } from './todo.service';
import { V1Todo } from '../../../protos/clients/ts/src/models';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, TodoListComponent, FormsModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  todos: V1Todo[] = [];
  newTodoText = '';
  title = `World's Best Angular-Go Todo App`;

  constructor(private todoService: TodoService) { }

  ngOnInit(): void {
    this.fetchTodos();
  }

  fetchTodos(): void {
    this.todoService.getTodos().subscribe((response: TodoApiResponse) => {
      this.todos = response.data;
    });
  }


  addTodo(): void {
    if (!this.newTodoText.trim()) {
      return; // Avoid adding empty todos
    }

    const newTodo: V1Todo = {
      todoName: this.newTodoText,
      todoType: 'TODO_GENERIC',
      priority: 1,
      completed: false,
    };

    console.log({ newTodo })

    this.todoService.addTodo(newTodo as V1Todo).subscribe({
      next: (response) => {
        console.log('Received response:', response);
        const receivedTodos = response.data;
        this.todos = [...this.todos, ...receivedTodos];
        this.newTodoText = '';
      },
      error: (error) => {
        console.error('Error adding todo', error);
      }
    });
  }
}
