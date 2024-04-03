import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import type { V1Todo } from '../../../protos/clients/ts/src/models'

@Injectable({
  providedIn: 'root'
})
export class TodoService {
  private apiUrl = 'http://localhost:6001/v1/todos'

  constructor(private http: HttpClient) { }

  getTodos(): Observable<V1Todo[]> {
    return this.http.get<V1Todo[]>(this.apiUrl + '/list');
  }

  addTodo(todo: V1Todo): Observable<{ data: V1Todo[] }> {
    const newTodo = { data: [todo] };
    return this.http.post<{ data: V1Todo[] }>(this.apiUrl, newTodo);
  }
}
