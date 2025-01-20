package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type StructA struct {
    FieldA string `form:"field_a"`
}

type StructB struct {
    NestedStructA StructA
    FieldB string `form:"field_b"`
}

type StructC struct {
    NestedStructA *StructA
    FieldC string `form:"field_c"`
}

type StructD struct {
    NestedStructX struct {
        FieldX string `form:"field_x"`
    }
    FieldD string `form:"field_d"`
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
            "a": b.NestedStructA,
            "b": b.FieldB,
        },
    )
}

func getC(context *gin.Context) {
    var c StructC
    context.Bind(&c)
    context.JSON(
        http.StatusOK,
        gin.H {
            "a": c.NestedStructA,
            "c": c.FieldC,
        },
    )
}

func getD(context *gin.Context) {
    var d StructD
    context.Bind(&d)
    context.JSON(
        http.StatusOK,
        gin.H {
            "x": d.NestedStructX,
            "d": d.FieldD,
        },
    )
}

func main() {
    router := gin.Default()

    router.GET("/", getRoot)
    router.GET("/get-a", getA)
    router.GET("/get-b", getB)
    router.GET("/get-c", getC)
    router.GET("/get-d", getD)

    router.Run(":8080")
}
