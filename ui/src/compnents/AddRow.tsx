import * as React from 'react'
import { Button,Input } from 'semantic-ui-react'
import { MessageService } from '../services/messageService';

interface IProps{
    refreshMessages: any
}

interface IState{
    text: string
}

export class AddRow extends React.Component<IProps, IState> {


    constructor(props:any){
        super(props)
        this.state = {text:''}
    }

     public render(){
         return(
            <div>
                 <Input fluid={true} value={this.state.text} onChange={this.handleChange} placeholder='Add A New Message' />
                 <Button onClick={this.createMessage} basic={true} color='purple' content='Add' />
            </div>
         )
     }

    private handleChange = ( event:any):any =>{
        this.setState({text: event.target.value})
    }

    private createMessage = ()=> {
        MessageService.createMessage(this.state.text).then(res =>  {
           this.setState({text: ''}) 
           this.props.refreshMessages();
        },
        err =>{
            // tslint:disable-next-line:no-console
            console.error("error")
        })
    }
}