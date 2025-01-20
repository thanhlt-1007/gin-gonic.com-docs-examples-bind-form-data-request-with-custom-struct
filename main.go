package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStruct StructA
    FieldB string `form:"field_b"`
}

func getRoot(context *gin.Context) {
    context.JSON(
        http.StatusOK,
        gin.H {
            "message": "success",
        },
    )
}

func getA(context *gin.Context) {
    var a StructA
    context.Bind(&a)
    context.JSON(
        http.StatusOK,
        gin.H{
            "a": a.FieldA,
        },
    )
}

func getB(context *gin.Context) {
    var b StructB
    context.Bind(&b)
    context.JSON(
        http.StatusOK,
        gin.H{
            "a": b.NestedStruct,
            "b": b.FieldB,
        },
    )
}

func main() {
    router := gin.Default()
    router.GET("/", getRoot)
    router.GET("/get-a", getA)
    router.GET("/get-b", getB)

    router.Run(":8080")
}
