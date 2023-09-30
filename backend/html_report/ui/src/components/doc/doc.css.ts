import { css } from "lit";

export default css`


:host {
    height: calc(100vh - 80px);
    position: relative;
    background-color: var(--gray22);
    border-top-right-radius: 10px;
    border-top-left-radius: 10px;
    display: block;
  }

  .labels {
    position: absolute;
    right: 0%;
    transform: translateX(50%);
    top: 5%;
    background-color: var(--gray30);
    padding: 1.2rem 3.2rem;
    font-size: 2rem;
    border-radius: 10px;
    color: var(--80);
    border: none;
  }

  .labels p {
    opacity: .5;
  }
  
  .labels .on {
    opacity: 1;
  }

  .doc {
    color: initial;
    background-color: white;
    width: 80%;
    margin: 0 auto;
    overflow: hidden;
    overflow-y: scroll;
    height: 100%;
  }
  
* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;

}

html {
    font-size: 62.5%;
}

.frame {
    max-width: 75%;
    background-color: white;
}

p {

}

`;