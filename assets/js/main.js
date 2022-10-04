
// url for api / retrieval
const URL = 'http://localhost:5000/'

$(window).bind("load", setUpTable());

setInterval(function () {
    var date = new Date()
    if (date.getMinutes() == 0){
        $.ajax({
            method: 'GET',
            url: URL + 'api/retrieve',
            success:function(response){
                console.log(response)
            }
        })
    }
}, 30 * 1000)

function setUpTable(){
    // sets up data first
    updateTable()

    // for click, retrieve data and sort
    $('th').on('click', function (){
        var tableHeader = this
        $.ajax({
            method: 'GET',
            url:'api/trips/gets-data-display',
            success:function(response){
                var data = response 
                
                var column = $(tableHeader).data('column')
                var order = $(tableHeader).data('order')

                if(order == 'desc'){
                    $(tableHeader).data('order', 'asc')
                    data = data.sort((a,b) => a[column] > b[column] ? 1 : -1)
                } else {
                    $(tableHeader).data('order', 'desc')
                    data = data.sort((a,b) => a[column] < b[column] ? 1 : -1)
                }

                setData(data)
            }
        })
    })
    
}

function updateTable(){
    $.ajax({
        method: 'GET',
        url:'http://localhost:5000/api/trips/gets-data-display',
        success:function(response){
            setData(response)
        }
    })
}

function setData(data){
    // update table
    var table = document.getElementById('dataTable')
    table.innerHTML = ''
    for (var i = 0; i < data.length; i++){
        var row = `<tr style="text-align: left;">
                        <td>${data[i].name}</td>
                        <td>${data[i].distance_traveled}</td>
                        <td>${data[i].co2_produced}</td>
                        <td>${data[i].plastic_straws_used}</td>
                    </tr>`
        table.innerHTML += row
    }
}

