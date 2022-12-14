<script>
</script>

<template>
    <div id="artCont">
        <article>
            <div id="imgCont">
                <div id="three">

                </div>
            </div>
        </article>
    </div>
</template>


<style>
#artCont {
    justify-content: flex-start;
    flex-direction: column;
    align-items: stretch;
    display: flex;
    position: relative;
    box-sizing: border-box;
    color: white;
}

article {
    flex-grow: 1;
    display: block;

}

#imgCont {
    position: relative;
    display: flex;
    flex-direction: column;
    padding-bottom: 0px;
    padding-top: 0px;
}

#three {
    align-items: stretch;
    box-sizing: border-box;
    display: flex;
    flex-direction: row;
    flex-shrink: 0;
    position: relative;
    margin-bottom: 28px;
}
</style>