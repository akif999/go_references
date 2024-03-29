# Go の go/ast package のデータ構造についてのまとめ

## 概要 Overview

Go の `go/ast` package のデータ構造についてまとめます。  
Summarize the data structure of Go's `go / ast` package.

## 目的

skeleton 等を使用した静的解析ツール/モジュールを作成する際に、  
`go/ast` package のデータ構造についての情報を参照しやすくすることを目的とします。  

## 背景

`go/ast` package のデータ構造は、ネストした Interface 型を持つ構造となっており、  
複雑性の高いものとなっています。  
それ故、公式のドキュメントや、ソースコードからデータ構造の全体像をつかむには、
コストがかります。  
よって、その全体像について、把握をしやすくすることを検討しました。  

## go/ast package の基本情報


## go/ast のデータ構造 Data structure of go/ast


|| No. || type || name || description || example || addition ||

## Go のソースコードの解析方法について

以下のようにいくつかの選択肢がある。

* ソースコードを Tokens に分解して順に捜査する
    * AST の作成前の段階のため、あくまで Token としての情報しか持っていない
* ソースコードを AST へ変換してから順に捜査する
    * ざっと File -> Decl -> Statement -> Expression の階層でツリーになっている

## 参照 References

* https://monpoke1.hatenablog.com/entry/2018/12/16/110943#ValueSpec
* https://mom0tomo.github.io/post/go_ast_parser_static_analysis_2/
