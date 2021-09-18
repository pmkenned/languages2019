"use strict";

const myEvents = [
    {
        'title': 'event1',
        'day': 'sunday',
        //'startTime': '8:15',
        //'endTime': '9:00',
        'startTime': '16:15',
        'endTime': '17:00',
        'color': 'red',
    },
    {
        'title': 'event2',
        'day': 'wednesday',
        //'startTime': '10:15',
        //'endTime': '10:30',
        'startTime': '17:15',
        'endTime': '17:30',
        'color': 'green',
    },
    {
        'title': 'event3',
        'day': 'tuesday',
        'startTime': '19:15',
        'endTime': '19:30',
        //'startTime': '13:00',
        //'endTime': '14:00',
        'color': 'blue',
    },
];

//const unitsPerHour = 4;
//const startHour = 8;
//const endHour = 20;
const unitsPerHour = 12;
const minutesPerUnit = Math.round(60/unitsPerHour)
const startHour = 16;
const endHour = 20;

let mouseDown = false;
let startSlot = null;
let endSlot = null;

function drawEvent(start_col, start_row, end_col, end_row, color) {
    for (let c = start_col; c <= end_col; c++) {
        if (c == start_col) {
            for (let r = start_row; r <= end_row; r++) {
                $(`#slot_${c}_${r}`).css('background-color', color);
            }
        }
    }
}

function mouse_down_func(e) {
    mouseDown = true;
    startSlot = e.target;
    endSlot = e.target;
}

function mouse_up_func(e) {
    mouseDown = false;
    endSlot = e.target;
    console.log('start slot: ',  startSlot.id, 'end slot: ',  endSlot.id);
    let startMatches = startSlot.id.match(/slot_(\d+)_(\d+)/);
    startMatches.shift();
    startMatches = startMatches.map((item) => parseInt(item));
    console.log(startMatches);
    let endMatches = endSlot.id.match(/slot_(\d+)_(\d+)/);
    endMatches.shift();
    endMatches = endMatches.map((item) => parseInt(item));
    console.log(endMatches);

    const start_hour_military = Math.floor(startMatches[1]/unitsPerHour) + startHour;
    const start_min = (startMatches[1] % unitsPerHour) * minutesPerUnit;
    const start_min_str = start_min.toString().padStart(2, "0")
    const startTime = `${start_hour_military}:${start_min_str}`;

    const end_hour_military = Math.floor(endMatches[1]/unitsPerHour) + startHour;
    const end_min = (endMatches[1] % unitsPerHour) * minutesPerUnit;
    const end_min_str = end_min.toString().padStart(2, "0")
    const endTime = `${end_hour_military}:${end_min_str}`;

    console.log('START TIME: ', startTime);
    console.log('END TIME: ', endTime);

    myEvents.push({'title': 'event4', 'day': 'friday', 'startTime' : startTime, 'endTime' : endTime, 'color': 'orange'});
    drawEvents();
}

window.addEventListener('mousemove', e => {
    if (mouseDown) {
        endSlot = e.target;
        //$(e.target).css('background-color', '#800000');
    }
});

let myNext = $('#sidebar');
myNext.after(`<div class="hour_col"></div>`);
myNext = myNext.next();
for(let i=0; i<(endHour-startHour)*unitsPerHour; i++) {
    const hour_military = Math.floor(i/unitsPerHour) + startHour;
    const hour = (hour_military <= 12) ? hour_military : hour_military - 12;
    const am_pm = (hour_military <= 12) ? 'am' : 'pm';
    const min = (i % unitsPerHour) * minutesPerUnit;
    const min_str = min.toString().padStart(2, "0")
    const content = (min == 0) ? `${hour}:${min_str} ${am_pm}` : '';
    myNext.after(`<div class="hour_col">${content}</div>`);
    myNext = myNext.next();
}

const myIDs = ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday'];
const myClasses = ['sun', 'mon', 'tue', 'wed', 'thu', 'fri', 'sat'];

for (let idx in myIDs) {
    const id = myIDs[idx];
    const _class = myClasses[idx];
    const day = $(`#${id}`);
    myNext = day;

    for(let i=0; i<(endHour - startHour)*unitsPerHour; i++) {
        //const content = `${_class} ${i}`;
        const content = '';
        myNext.after(`<div class="${_class} slot">${content}</div>`);
        myNext = myNext.next();
        myNext.attr('id', `slot_${idx}_${i}`);
        myNext.mousedown((e) => mouse_down_func(e));
        myNext.mouseup((e) => mouse_up_func(e));
    }
}

function drawEvents() {
    for (let e of myEvents) {
        let startHourMinute = e.startTime.split(':').map((item) => parseInt(item));
        let endHourMinute = e.endTime.split(':').map((item) => parseInt(item));
        let eventStartHour = startHourMinute[0];
        let eventStartMinutes = startHourMinute[1];
        let eventEndHour = endHourMinute[0];
        let eventEndMinutes = endHourMinute[1];
        //console.log(eventStartHour, eventStartMinutes, eventEndHour, eventEndMinutes);
        const days = ['sunday', 'monday', 'tuesday', 'wednesday', 'thursday', 'friday', 'saturday'];
        const start_col = days.indexOf(e.day);
        const start_row = (eventStartHour - startHour)*unitsPerHour + Math.floor(eventStartMinutes/minutesPerUnit);
        const end_col = start_col;
        const end_row = (eventEndHour - startHour)*unitsPerHour + Math.floor(eventEndMinutes/minutesPerUnit) -1;
        //console.log(start_col, start_row, end_col, end_row);
        drawEvent(start_col, start_row, end_col, end_row, e.color);
    }
}

drawEvents();
