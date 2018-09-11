import * as React from 'react';
import { Container, Segment } from 'semantic-ui-react';
import './App.css';
import { AddRow } from './compnents/AddRow';
import { MessageList } from './compnents/MessageList';
import { IMessage } from './models';
import { MessageService } from './services/messageService';

class App extends React.Component<{},{messageList:IMessage[]}> {
  constructor(props: any) {
    super(props)
    this.state = { messageList: []}
    MessageService.getMessages().then(messageList=>{
      this.setState({messageList})
    })
  }

  public refreshMessages(){
    MessageService.getMessages().then(messageList=>{
      this.setState({messageList})
    })
  }

  public render() {
    const refreshMethod = this.refreshMessages.bind(this)
    return (
      <Container>
        <header className="App-header">
          <h1 className="App-title">PalinDetect</h1>
        </header>
        <Segment color='yellow' inverted={true}>
        <Container fluid={true}>
        <AddRow refreshMessages={refreshMethod}/>
        <MessageList refreshMessages={refreshMethod} messages={this.state.messageList}/>
        </Container>
        </Segment>
      </Container>
    );
  }
}

export default App;
