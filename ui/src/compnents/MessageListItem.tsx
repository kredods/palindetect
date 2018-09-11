import * as React from 'react';
import { Button, Dropdown,Header,Icon, Input, Modal, Segment } from 'semantic-ui-react';
import { IMessage } from "../models";
import { MessageService } from '../services/messageService';
interface IProps{
    message: IMessage
    refreshMessages: any
    key: any
}

interface IState{
    isEdit: boolean
    updatedMessage: string
}

export class MessageListItem extends React.Component<IProps, IState>{

    constructor(props: any){
        super(props)
        this.state = {
                isEdit: false,
                updatedMessage: props.message.body
            }
    }

    public render(){
        return (
            <Segment attached={true} fluid={true} inverted={true} color='purple'>
                {!this.state.isEdit?(
                <div>
               <p className='message-text'>{this.props.message.body}</p>
               <Dropdown text='...'>
                    <Dropdown.Menu>
                        <Modal trigger={<Dropdown.Item text='Details' />} basic={true} size='small' closeIcon={true}>
                        <Header icon='comment' content={`${this.props.message.body} is ${this.props.message.isPalindrome? '' :'not' } a palindrome` } />
                    </Modal>
                        <Dropdown.Item text='Edit' onClick={this.setEdit} />
                        <Dropdown.Item onClick={this.deleteMessage}  text='Delete'/>
                    </Dropdown.Menu>
                </Dropdown>
                </div>
                ):(
                    <div>
                    <Input onChange={this.updateMessage} value ={this.state.updatedMessage} placeholder='Enter new string'/>
                    <Button onClick={this.clearEdit} color='red' inverted={true}>
                        <Icon name='remove' /> Clear
                     </Button>
                    <Button onClick={this.saveMessage} color='green' inverted={true}>
                        <Icon name='checkmark' /> Save
                    </Button>
                    </div>

                )}
            </Segment>
        )
    }

    public updateMessage = (event: any) =>{
        this.setState(
            {
                isEdit: this.state.isEdit,
                updatedMessage: event.target.value
            }
        )
    }

   public setEdit = () => {
       this.setState({
           isEdit: true,
           updatedMessage: this.state.updatedMessage 
        })
   }

   public clearEdit = () => {
    this.setState({
        isEdit: false,
        updatedMessage: this.state.updatedMessage 
     })
   }


   public saveMessage = () => {
       MessageService.updateMessage(this.props.message.id, this.state.updatedMessage ).then( () =>{
        this.props.refreshMessages()
        this.clearEdit()
    },
    err =>{
            // tslint:disable-next-line:no-console
            console.error("failed to update");
            this.clearEdit()
    })
   }

    public deleteMessage = () => {
        MessageService.deleteMessage(this.props.message.id).then( () =>{
            this.props.refreshMessages()
        },
        err =>{
                // tslint:disable-next-line:no-console
                console.error("failed to delete");
        })
    }   
}
