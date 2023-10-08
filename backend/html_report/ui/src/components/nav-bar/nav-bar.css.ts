import { css } from "lit";

export default css`

* {
    margin: 0;
    padding: 0;
    box-sizing: border-box;

}

html {
    font-size: 62.5%;
}

nav {
    display: flex;
    background-color: var(--gray20);
    justify-content: space-between;
    align-items: center;
    padding: 0 2.4rem;
    font-size: 2.4rem;
    height: 70px; 
}

ul {
    align-self: end;
    display: flex;
    align-items: end;
    gap: .6rem;
    list-style: none;
    // height: 100%;
    position: relative;
}

h3, .tab {
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
    background-color: var(--gray22);
    padding: 1.2rem 2.4rem;
    align-self: end;
    color: var(--gray60);
}

.tab {
    display: block;
    border-bottom: 2px solid var(--gray25);
    height: 100%;
    cursor: pointer;
    transition: all .2s;
}

.selected-tab {
    color: var(--gray);
    border-bottom: 2px solid var(--info);
}
h3 {
    display: flex;
    align-items: center;
    gap: .6rem;
    justify-content: center;
}

input {
    background-color: transparent;
    border: none;
    padding: .6rem;
    color: var(--gray80);
    font-size: 2.4rem;
    width: 12rem;
    text-align: center; 
}

input {
    font-family: NeueMachina !important;
  }
  

.fill-in-name {
    box-shadow: 0 0 0 1px var(--gray92);
    border-radius: 10px;    
}

.popup {
    position: absolute;
    visibility: hidden;
    opacity: 0;
    // display: none;
    top: 50%;
    transform: translate(100%, -50%);
    right: 0;
    transition: all .2s;
}

.brighten {
    opacity: 1 !important;
    visibility: visible; 
}
`;