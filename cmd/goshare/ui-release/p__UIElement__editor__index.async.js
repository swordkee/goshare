(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[5],{"4zw2":function(e,t,n){"use strict";n.r(t),n.d(t,"default",function(){return j});n("14J3");var a=n("BMrR"),r=(n("jCWc"),n("kPKH")),o=(n("IzEo"),n("bx4M")),i=n("2Taf"),l=n.n(i),c=n("vZ4D"),d=n.n(c),s=n("l4Ni"),u=n.n(s),m=n("ujKo"),E=n.n(m),g=n("MhPg"),v=n.n(g),C=n("q1tI"),b=n.n(C),h=n("Kvkj"),w=n("v83y"),f=n("bQ8i"),p=n.n(f),y=n("xjsL"),S=n.n(y),j=function(e){function t(e){var n;return l()(this,t),n=u()(this,E()(t).call(this,e)),n.onEditorStateChange=function(e){n.setState({editorContent:e})},n.state={editorContent:null},n}return v()(t,e),d()(t,[{key:"render",value:function(){var e=this.state.editorContent,t={lg:12,md:24,style:{marginBottom:32}},n={minHeight:496,width:"100%",background:"#f7f7f7",borderColor:"#F1F1F1",padding:"16px 8px"};return b.a.createElement("div",{className:"content-inner"},b.a.createElement(a["a"],{gutter:32},b.a.createElement(r["a"],t,b.a.createElement(o["a"],{title:"Editor",style:{overflow:"visible"}},b.a.createElement(h["b"],{wrapperStyle:{minHeight:500},editorStyle:{minHeight:376},editorState:e,onEditorStateChange:this.onEditorStateChange}))),b.a.createElement(r["a"],t,b.a.createElement(o["a"],{title:"HTML"},b.a.createElement("textarea",{style:n,disabled:!0,value:e?p()(Object(w["convertToRaw"])(e.getCurrentContent())):""}))),b.a.createElement(r["a"],t,b.a.createElement(o["a"],{title:"Markdown"},b.a.createElement("textarea",{style:n,disabled:!0,value:e?S()(Object(w["convertToRaw"])(e.getCurrentContent())):""}))),b.a.createElement(r["a"],t,b.a.createElement(o["a"],{title:"JSON"},b.a.createElement("textarea",{style:n,disabled:!0,value:e?JSON.stringify(Object(w["convertToRaw"])(e.getCurrentContent())):""})))))}}]),t}(b.a.Component)}}]);