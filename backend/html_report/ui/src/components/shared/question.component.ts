import { LitElement, html } from 'lit'
import { customElement, property } from 'lit/decorators.js'
import questionCss from './question.css';


// 3 Steps


@customElement('question-component')
export class Question extends LitElement {

  static styles = [questionCss]

  @property()
  description: string ="";
 
  @property()
  width: number = 0;

  constructor() {
    super()
}


render(){
    return html`
    ?
        <span
        style="width: ${this.width + "rem"};"
        class="explanation"
        >${this.description}</span>
    `
  }
}
