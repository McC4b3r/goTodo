import { Component, OnInit } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { FormsModule } from '@angular/forms';
import { TodoListComponent } from './components/todo-list/todo-list.component';
import { TodoService } from './todo.service';
import { V1Todo } from '../../../protos/clients/ts/src/models';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, TodoListComponent, FormsModule],
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  todos: V1Todo[] | undefined;
  newTodoText = '';
  title = `World's Best Angular-Go Todo App`;

  constructor(private todoService: TodoService) { }

  ngOnInit(): void {
    this.fetchTodos();
  }

  fetchTodos(): void {
    this.todoService.getTodos().subscribe((todos) => {
      this.todos = todos;
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

    this.todoService.addTodo(newTodo).subscribe({
      next: (response) => {
        console.log('Added new todo:', response.data);
        this.todos = this.todos ? [...this.todos, ...response.data] : [...response.data];
        this.newTodoText = ''; // Clear the input field
      },
      error: (error) => {
        console.error('Error adding todo', error);
      }
    });
  }



}
