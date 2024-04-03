<<<<<<< HEAD
import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import type { V1Todo } from '../../../protos/clients/ts/src/models'
=======
import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';
import { V1Todo } from '../../../protos/clients/ts/src/models/V1Todo';
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a

@Injectable({
  providedIn: 'root'
})
export class TodoService {
<<<<<<< HEAD
  private apiUrl = 'http://localhost:6001/v1/todos'
=======
  private apiUrl = 'http://localhost:6001/v1/todos/list'
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a

  constructor(private http: HttpClient) { }

  getTodos(): Observable<V1Todo[]> {
<<<<<<< HEAD
    return this.http.get<V1Todo[]>(this.apiUrl + '/list');
  }

  addTodo(todo: V1Todo): Observable<{ data: V1Todo[] }> {
    const newTodo = { data: [todo] };
    return this.http.post<{ data: V1Todo[] }>(this.apiUrl, newTodo);
=======
    return this.http.get<V1Todo[]>(this.apiUrl);
  }

  postTodo(todo: V1Todo): Observable<V1Todo> {
    return this.http.post<V1Todo>(this.apiUrl, todo);
  }

  deleteTodo(id: number): Observable<V1Todo> {
    return this.http.delete<V1Todo>(`${this.apiUrl}/${id}`);
  }

  updateTodo(todo: V1Todo): Observable<V1Todo> {
    return this.http.put<V1Todo>(`${this.apiUrl}/${todo.id}`, todo);
>>>>>>> dd2550794e41b87870bfb030aafc166cd14fdd3a
  }
}
