import { IMessage } from '../models/message';
const baseHost = "http://13.57.28.144:8000" || "http://localhost:8000"

export class MessageService {
public static createMessage(message: string){
    const url = `${baseHost}/v1/messages`
    return fetch(url, {
        body: JSON.stringify({
            body: message
        }),
        headers: {
            'Content-Type': 'application/json',
          },
        method: 'POST',
      }).then(response => response.json())
    }

public static updateMessage(id:any, update: string){
    const url = `${baseHost}/v1/messages/${id}`
    return fetch(url, {
    body: JSON.stringify({
        body: update,
        id
    }),
    headers: {
        'Content-Type': 'application/json',
        },  
    method: 'PUT', 
    }).then(response => response.json())
}


public static deleteMessage(id: any): Promise<IMessage>{
    const url = `${baseHost}/v1/messages/${id}`
    return fetch(url, {
    headers: {
        'Content-Type': 'application/json',
        },  
    method: 'DELETE', 
    }).then(response => response.json())
}

  public static getMessages(): Promise<IMessage[]> {
    const url = `${baseHost}/v1/messages`
  return fetch(url)
    .then(response => response.json())
    .then(MessageService.extractMessages)
  }

  private static extractMessages(messages: any[]) :IMessage[] {
   if(!messages){
       return []
   }  
   return messages.map(MessageService.extractMessage)
  }

  private static extractMessage(message: any): IMessage{
      return {
          body: message.body,
          id: message.id,
          isPalindrome: message.isPalindrome
      }
  }


}