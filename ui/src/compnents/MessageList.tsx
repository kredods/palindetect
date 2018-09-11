import * as React from 'react'
import { Segment} from 'semantic-ui-react'
import { IMessage } from '../models/message';
import { MessageListItem } from './MessageListItem';

interface IProps {
   messages: IMessage[]
   refreshMessages: any
}

interface IState{
    messages: IMessage[]
}
export class MessageList extends React.Component<IProps,IState>{
    constructor(props: any){
        super(props)
        this.state = this.props
    }

    public render(){
        return (
            <Segment.Group>
            {
                this.props.messages.length !==0 ?
                this.props.messages.map( message =>(
                    <MessageListItem refreshMessages={this.props.refreshMessages}  message={message}/>
                ))
                : (
                    <p>No Messages Created</p>
                )
            }
            </Segment.Group>
        )
    }
}