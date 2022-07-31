/*
*                  _________________
*                  `.    terminal   `.         interface(CreateRobot)
*                    `––––––––––––––––`
*                  /           |         \
*  _____________      ______________       ________________
*  `. WecomRobot`.     `.  LarkRobot `.     `. DingDingRobot`.       RobotFactory
*    ` –––––––––––`      `–––––––––––––`      `—————————————— `
*          |                    |                    |
*          |                    |                    |
*   _____________       ______________        _______________
*   `.    Wecom  `.      `.   Larkshu `.      `.   DingDing  `.      Robot
*     ` –––––––––––`       `––––––––––––`       `——————————————`
*                 \             |                /
*                  \            |               /
*                   \     ________________     /
*                         `.  RobotBehavior`.         interface(SendSpecifiedAddress)
*                           `––––––––––––––––`
*
 */

package output

import "papillon/event"

//terminal is RobotFactory
type Terminal interface {
	CreateRobot(msg *event.MessageNotify) Robot
}

//Robot defines some behaviors
type Robot interface {
	SendSpecifiedAddress() error
}
