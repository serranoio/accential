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

:host {
  height: calc(100vh - 80px);
  position: relative;
  display: block;
  background-color: var(--gray22);
  border-top-right-radius: 10px;
  border-top-left-radius: 10px;
  overflow: hidden;
  overflow-y: scroll;
}

.row {
    display: flex;
    border-right: 1px solid var(--gray25);
    border-left: 1px solid var(--gray25);
    font-size: 2rem;
    position: relative;
    justify-content: space-between;
  }
  
  .titles {
    color: var(--gray92);
  }

  .second {
    
    color: var(--gray80);
  }

  .column {
    display: block;
    border: 1px solid var(--gray25);
    display: flex;
    align-items: center;
    padding: var(--spacingHalf) var(--spacingHalf); 
    font-size: 2rem;
    flex: 1;

  }
  
  .report {
    display: flex;
    justify-content: center;
    width: 100%;
  }

  tbody {
    width: 100%;
  }
  

  .include-metric,
  .more, .edit-metric {
    padding: var(--spacingQuarter) var(--spacingHalf);
    border-radius: 10px;
    background: none;
    border: none;
    font-size: 2rem;
    cursor: pointer;
    box-shadow: 0 0 0 2px var(--gray60);
    color: var(--gray60);
    transition: all .2s;
  }
  
  .edit-metric:hover,
  .include-metric:hover,
  .more:hover {
    box-shadow: 0 0 0 2px var(--gray80);
    color: var(--gray80);
    
  }
  
  .include-metric {
    margin-right: 2.4rem;
    box-shadow: 0 0 0 2px var(--success);
    color: var(--success);
  }
  
  .include-metric:hover {
    box-shadow: 0 0 0 2px var(--successD10);
    color: var(--successD10);

  }

  
  .info {
    overflow: hidden;
    // transition: all .5s;
    transform-origin: top;
    height: 0;
    visibility: hidden;
    min-height: 0px;
  }
  .show {
    overflow: visible;
    padding: var(--spacingHalf);
    visibility: visible;
    min-height: 0px;
    height: auto;
  }

  .turn {
    transform: rotate(180deg);
  }

  
`;